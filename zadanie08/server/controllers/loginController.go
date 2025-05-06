package controllers

import (
	"errors"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"zadanie08/database"
	"zadanie08/models"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateSession(c echo.Context, user models.User) error {
	sess, _ := session.Get("session", c)

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400, // 1 dzień
		HttpOnly: true,
	}

	sess.Values["email"] = user.Email
	sess.Values["name"] = user.Name
	sess.Values["surname"] = user.Surname
	err := sess.Save(c.Request(), c.Response())
	if err != nil {
		return err
	}

	return nil
}

func LoginUser(c echo.Context) error {
	var loginRequest LoginRequest
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload: " + err.Error()})
	}

	var user models.User
	result := database.DB.Where("email = ?", loginRequest.Email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve user: " + result.Error.Error()})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	err = CreateSession(c, user)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Zalogowano i ustawiono sesję.")
}

func Logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Values = map[interface{}]interface{}{}
	err := sess.Save(c.Request(), c.Response())
	if err != nil {
		fmt.Println("Failed to destroy session")
		return err
	}
	return c.String(http.StatusOK, "Wylogowano.")
}
