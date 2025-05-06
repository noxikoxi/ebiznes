package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

	e := echo.New()

	routers.RegisterRoutes(e)
	routers.LoginRoutes(e)
	
	e.Logger.Fatal(e.Start(":1323"))
}
