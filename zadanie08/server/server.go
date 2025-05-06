package main

import (
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"zadanie08/database"
	"zadanie08/models"

	"github.com/labstack/echo/v4"
	"zadanie08/routers"
)

func main() {

	db, err := gorm.Open(sqlite.Open("database/database.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	db.Exec("PRAGMA foreign_keys = ON")
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("Failed to migrate database")
		return
	}
	database.InitDB(db)
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowOrigins: []string{"*"},
	}))

	routers.RegisterRoutes(e)
	routers.LoginRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
