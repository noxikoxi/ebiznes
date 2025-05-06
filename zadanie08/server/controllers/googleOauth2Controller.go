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
	"zadanie08/configs"
	"zadanie08/database"
	"zadanie08/models"
	"zadanie08/utils"
)

var oauthStateString = "randomString"

// generuje url do logowania do google
func HandleGoogleLogin(c echo.Context) error {
	url := configs.GoogleOauth2Config.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
	return c.JSON(http.StatusOK, url)
}

// odbiera token od google
func HandleGoogleCallback(c echo.Context) error {
	code := c.QueryParam("code")
	if code == "" {
		return c.String(http.StatusBadRequest, "Code not found")
	}

	// Wymiana kodu na token dostępu
	token, err := configs.GoogleOauth2Config.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Error exchanging code for token: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to exchange code for token")
	}

	client := configs.GoogleOauth2Config.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Printf("Error getting user info: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to get user info")
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		log.Printf("Error decoding user info: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to decode user info")
	}

	email, _ := userInfo["email"].(string)
	givenName, _ := userInfo["given_name"].(string)
	familyName, _ := userInfo["family_name"].(string)

	var existingUser models.User
	result := database.DB.Where("email = ?", email).First(&existingUser)

	if result.Error == nil {
		database.DB.Model(&existingUser.Tokens).Update("GoogleAccessToken", token.AccessToken)
		database.DB.Model(&existingUser.Tokens).Update("GoogleTokenExpires", token.Expiry)

		err := CreateSession(c, existingUser)
		if err != nil {
			return err
		}

		jwtToken, err := utils.GenerateJWT(existingUser.ID, existingUser.Email, existingUser.Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate JWT"})
		}

		database.DB.Model(&existingUser.Tokens).Update("JWT", jwtToken)

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
			Email:    email,
			Name:     givenName,
			Surname:  familyName,
			Password: "",
		}
		result := database.DB.Create(&newUser)
		if result.Error != nil {
			log.Printf("Błąd podczas tworzenia użytkownika: %v", result.Error)
			return c.String(http.StatusInternalServerError, "Nie udało się utworzyć użytkownika")
		}

		jwtToken, err := utils.GenerateJWT(existingUser.ID, existingUser.Email, existingUser.Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate JWT"})
		}
		// Optionally create tokens record
		database.DB.Create(&models.Tokens{
			UserID:             newUser.ID,
			GoogleAccessToken:  token.AccessToken,
			GoogleTokenExpires: token.Expiry,
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
