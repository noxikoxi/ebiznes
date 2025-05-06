package controllers

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetUserData(c echo.Context) error {
	sess, _ := session.Get("session", c)
	
	email := sess.Values["email"]
	name := sess.Values["name"]
	surname := sess.Values["surname"]

	if email == nil {
		return c.String(http.StatusUnauthorized, "Brak sesji lub niezalogowany użytkownik.")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"email":   email,
		"name":    name,
		"surname": surname,
	})
}
