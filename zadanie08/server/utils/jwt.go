package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint, email, name string) (string, error) {
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	expirationTime := time.Now().Add(24 * time.Hour) // Token ważny przez 24 godziny
	claims := &Claims{
		UserID: userID,
		Email:  email,
		Name:   name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
