package controllers

import (
	"errors"
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
	return c.JSON(http.StatusOK, map[string]interface{}{
		"email":   user.Email,
		"name":    user.Name,
		"surname": user.Surname,
	})
}
