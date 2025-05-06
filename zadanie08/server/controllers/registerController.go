package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"zadanie08/database"
	"zadanie08/models"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type UserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Surname  string `json:"surname" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

func RegisterUser(c echo.Context) error {
	var userRequest UserRequest

	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload: " + err.Error()})
	}

	validate := validator.New()
	if err := validate.Struct(userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Validation error: " + err.Error()})
	}

	var existingUser models.User
	result := database.DB.Where("email = ?", userRequest.Email).First(&existingUser)
	if result.Error == nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": "Email already used"})
	}

	hashedPassword, err := hashPassword(userRequest.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password: " + err.Error()})
	}

	newUser := models.User{
		Email:    userRequest.Email,
		Name:     userRequest.Name,
		Surname:  userRequest.Surname,
		Password: hashedPassword,
	}
	result = database.DB.Create(&newUser)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user: " + result.Error.Error()})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"email":   newUser.Email,
		"name":    newUser.Name,
		"surname": newUser.Surname,
	})
}
