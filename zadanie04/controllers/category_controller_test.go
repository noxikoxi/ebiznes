package controllers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"zadanie04/database"
	"zadanie04/models"
)

func setupDatabaseCategories() {
	// Tworzenie bazy w pamięci
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("could not connect to in-memory database")
	}

	// Migracje, żeby tworzyć tabele
	db.AutoMigrate(&models.Product{})

	db.Create(&models.Category{ID: 1, Name: "Category1"})
	db.Create(&models.Category{ID: 2, Name: "Category2"})
	db.Create(&models.Category{ID: 3, Name: "Category3"})

	database.DB = db
}

func TestGetCategories(t *testing.T) {
	setupDatabaseCategories()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/categories", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := GetCategories(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	var categories []CategoryResponse
	err = json.Unmarshal(rec.Body.Bytes(), &categories)
	assert.NoError(t, err)
	expectedCount := 3
	assert.Equal(t, expectedCount, len(categories))
	assert.Equal(t, "Category1", categories[0].Name)
	assert.Equal(t, "Category2", categories[1].Name)
	assert.Equal(t, "Category3", categories[2].Name)
}

func TestGetCategory(t *testing.T) {
	setupDatabaseCategories()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/categories/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	err := GetCategory(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	var category CategoryResponse
	err = json.Unmarshal(rec.Body.Bytes(), &category)
	assert.NoError(t, err)
	assert.Equal(t, "Category1", category.Name)
}

func TestCreateCategory(t *testing.T) {
	setupDatabaseCategories()
	e := echo.New()
	categoryData := `{
		"name": "NewSuperCategory"
	}`
	req := httptest.NewRequest(http.MethodPost, "/categories", strings.NewReader(categoryData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := CreateCategory(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	var response CategoryResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "NewSuperCategory", response.Name)

	var category models.Category
	result := database.DB.First(&category, "name = ?", "NewSuperCategory")
	assert.NoError(t, result.Error)
}

func TestUpdateCategory(t *testing.T) {
	setupDatabaseCategories()
	categoryData := `{
		"name": "NewCategory"
	}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/categories/1", strings.NewReader(categoryData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	err := UpdateCategory(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	var response CategoryResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "NewCategory", response.Name)

	var category models.Category
	database.DB.First(&category, "id = ?", 1)
	assert.Equal(t, "NewCategory", category.Name)
}

func TestDeleteCategory(t *testing.T) {
	setupDatabaseCategories()
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/categories", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	err := DeleteCategory(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)

	var category models.Category
	result := database.DB.First(&category, "id=?", 1)
	assert.Error(t, result.Error)
}
