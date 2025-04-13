package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"zadanie04/database"
	"zadanie04/models"
	"zadanie04/routers"
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

	// Migrate the schema
	err = db.AutoMigrate(&models.Product{}, &models.Category{}, &models.CartItem{}, &models.Cart{})
	if err != nil {
		fmt.Println("Failed to migrate database")
		return
	}

	e := echo.New()
	database.InitDB(db)
	routers.RegisterProductRoutes(e)
	routers.RegisterCategoryRoutes(e)
	routers.RegisterCartRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
