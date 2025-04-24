package main

import (
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"productsServer/database"
	"productsServer/models"
	"productsServer/routers"
	"regexp"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := gorm.Open(sqlite.Open("database/database.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	// Ważne, bez tego nie sprawdza kluczy
	db.Exec("PRAGMA foreign_keys = ON")

	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.Product{}, &models.CartItem{}, &models.Cart{}, &models.Payment{})
	if err != nil {
		fmt.Println("Failed to migrate database")
		return
	}

	allowedOriginPattern := `^https?://[a-z0-9]+\.ngrok-free\.app$`
	re := regexp.MustCompile(allowedOriginPattern)

	e := echo.New()
	database.InitDB(db)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowOriginFunc: func(origin string) (bool, error) {
			if origin == "http://localhost:5137" {
				return true, nil
			}
			return re.MatchString(origin), nil
		},
	}))

	routers.RegisterProductRoutes(e)
	routers.RegisterPaymentCartRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
