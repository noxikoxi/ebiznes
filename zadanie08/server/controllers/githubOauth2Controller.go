package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
	"time"
	"zadanie08/configs"
	"zadanie08/database"
	"zadanie08/models"
	"zadanie08/utils"
)

var oauthStateStringGithub = "randomStringGithub"

// generuje url do logowania do google
func HandleGithubLogin(c echo.Context) error {
	url := configs.GithubOauth2Config.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
	return c.JSON(http.StatusOK, url)
}

func HandleGithubCallback(c echo.Context) error {
	code := c.QueryParam("code")
	if code == "" {
		return c.String(http.StatusBadRequest, "Code not found")
	}

	// Wymiana kodu na token dostępu GitHub
	token, err := configs.GithubOauth2Config.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Error exchanging code for GitHub token: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to exchange code for GitHub token")
	}
	client := configs.GithubOauth2Config.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		log.Printf("Error getting GitHub user info: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to get GitHub user info")
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		log.Printf("Error decoding GitHub user info: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to decode GitHub user info")
	}

	name, _ := userInfo["name"].(string)
	parts := strings.Split(name, " ")
	surname := ""
	if len(parts) >= 2 {
		surname = strings.Join(parts[1:], " ")
	}
	resp, err = client.Get("https://api.github.com/user/emails")
	if err != nil {
		log.Println("Email fetch error:", err)
		return c.String(http.StatusInternalServerError, "Failed to get GitHub user emails")
	}
	defer resp.Body.Close()

	var emails []struct {
		Email    string `json:"email"`
		Primary  bool   `json:"primary"`
		Verified bool   `json:"verified"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&emails); err != nil {
		log.Println("Email parse error:", err)
		return c.String(http.StatusInternalServerError, "Failed to decode GitHub user emails")
	}

	var primaryEmail string
	for _, emailInfo := range emails {
		if emailInfo.Primary {
			primaryEmail = emailInfo.Email
			break
		}
	}

	if primaryEmail == "" {
		primaryEmail = emails[0].Email
	}

	var existingUser models.User
	result := database.DB.Preload("Tokens").Where("email = ?", primaryEmail).First(&existingUser)

	if result.Error == nil {
		err := CreateSession(c, existingUser)
		if err != nil {
			return err
		}

		jwtToken, err := utils.GenerateJWT(existingUser.ID, existingUser.Email, existingUser.Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate JWT"})
		}

		database.DB.Model(&existingUser.Tokens).Updates(map[string]interface{}{"JWT": jwtToken, "GithubAccessToken": token.AccessToken, "GithubTokenExpires": time.Now().Add(8 * time.Hour)})

		sess, _ := session.Get("session", c)
		sess.Values["authToken"] = jwtToken
		err = sess.Save(c.Request(), c.Response())
		if err != nil {
			return err
		}
		fmt.Println("Zalogowano istniejącego użytkownika:", existingUser.Email)
		return c.Redirect(http.StatusFound, "http://localhost:5173/hello")
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// User does not exist, create a new one
		newUser := models.User{
			Email:    primaryEmail,
			Name:     name,
			Surname:  surname,
			Password: "",
		}
		result := database.DB.Create(&newUser)
		if result.Error != nil {
			log.Printf("Błąd podczas tworzenia użytkownika: %v", result.Error)
			return c.String(http.StatusInternalServerError, "Failed to create user")
		}

		jwtToken, err := utils.GenerateJWT(existingUser.ID, existingUser.Email, existingUser.Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate JWT"})
		}
		// Optionally create tokens record
		database.DB.Create(&models.Tokens{
			UserID:             newUser.ID,
			GithubAccessToken:  token.AccessToken,
			GithubTokenExpires: time.Now().Add(8 * time.Hour),
			JWT:                jwtToken,
		})

		err = CreateSession(c, newUser)
		if err != nil {
			return err
		}

		sess, _ := session.Get("session", c)
		sess.Values["authToken"] = jwtToken
		err = sess.Save(c.Request(), c.Response())
		if err != nil {
			return err
		}

		// Create session or return JWT for newUser
		fmt.Println("Utworzono i zalogowano nowego użytkownika:", newUser.Email)
		return c.Redirect(http.StatusFound, "http://localhost:5173/hello")
	} else {
		// Some other database error
		log.Printf("Błąd bazy danych: %v", result.Error)
		return c.String(http.StatusInternalServerError, "Błąd bazy danych podczas logowania")
	}
}
