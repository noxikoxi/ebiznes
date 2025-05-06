package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"os"
	"zadanie08/configs"
	"zadanie08/database"
	"zadanie08/models"

	"github.com/labstack/echo/v4"
	"zadanie08/routers"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Błąd podczas ładowania pliku .env:", err)
	}
}

func main() {
	db, err := gorm.Open(sqlite.Open("database/database.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	db.Exec("PRAGMA foreign_keys = ON")
	err = db.AutoMigrate(&models.User{}, &models.Tokens{})
	if err != nil {
		fmt.Println("Failed to migrate database")
		return
	}
	database.InitDB(db)
	configs.InitAuthConfig()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:5173"},
	}))

	e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))))

	routers.RegisterRoutes(e)
	routers.LoginRoutes(e)
	routers.UserRoutes(e)
	routers.GoogleRoutes(e)
	routers.GithubRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
