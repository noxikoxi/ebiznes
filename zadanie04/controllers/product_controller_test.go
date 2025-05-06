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

func setupDatabaseProducts() {
	// Tworzenie bazy w pamięci
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("could not connect to in-memory database")
	}

	// Migracje, żeby tworzyć tabele
	db.AutoMigrate(&models.Product{}, &models.Category{})

	db.Create(&models.Category{ID: 1, Name: "Category1"})
	db.Create(&models.Category{ID: 2, Name: "Category1"})
	db.Create(&models.Category{ID: 3, Name: "Category1"})

	db.Create(&models.Product{ID: 1, Name: "Item1", Price: 10, CategoryID: 1})
	db.Create(&models.Product{ID: 2, Name: "Item2", Price: 100, CategoryID: 1})
	db.Create(&models.Product{ID: 3, Name: "Item3", Price: 1000, CategoryID: 2})
	db.Create(&models.Product{ID: 4, Name: "Item4", Price: 10, CategoryID: 2})
	db.Create(&models.Product{ID: 5, Name: "Item5", Price: 10, CategoryID: 3})
	db.Create(&models.Product{ID: 6, Name: "Item6", Price: 1000, CategoryID: 3})

	database.DB = db
}

func TestGetProducts(t *testing.T) {
	setupDatabaseProducts()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := GetProducts(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	var products []ProductResponse
	err = json.Unmarshal(rec.Body.Bytes(), &products)
	assert.NoError(t, err)
	expectedCount := 6
	assert.Equal(t, expectedCount, len(products))
	assert.Equal(t, "Item1", products[0].Name)
	assert.Equal(t, float32(10.0), products[0].Price)
	assert.Equal(t, "Category1", products[0].Category)
}

func TestGetProductsFilteredPrice(t *testing.T) {
	setupDatabaseProducts()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/products?minPrice=101", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := GetProducts(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	var products []ProductResponse
	err = json.Unmarshal(rec.Body.Bytes(), &products)
	assert.NoError(t, err)
	expectedCount := 2
	assert.Equal(t, expectedCount, len(products))
	assert.Equal(t, "Item3", products[0].Name)
	assert.Equal(t, "Item6", products[1].Name)
}

func TestGetProductsFilteredCategory(t *testing.T) {
	setupDatabaseProducts()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/products?categoryID=1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := GetProducts(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	var products []ProductResponse
	err = json.Unmarshal(rec.Body.Bytes(), &products)
	assert.NoError(t, err)
	expectedCount := 2
	assert.Equal(t, expectedCount, len(products))
	assert.Equal(t, "Item1", products[0].Name)
	assert.Equal(t, "Item2", products[1].Name)
}
func TestGetProductsFilteredCategoryAndPrice(t *testing.T) {
	setupDatabaseProducts()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/products?categoryID=1&minPrice=11", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := GetProducts(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	var products []ProductResponse
	err = json.Unmarshal(rec.Body.Bytes(), &products)
	assert.NoError(t, err)
	expectedCount := 1
	assert.Equal(t, expectedCount, len(products))
	assert.Equal(t, "Item2", products[0].Name)
}

func TestGetProduct(t *testing.T) {
	setupDatabaseProducts()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	err := GetProduct(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	var product ProductResponse
	err = json.Unmarshal(rec.Body.Bytes(), &product)
	assert.NoError(t, err)
	assert.Equal(t, "Item1", product.Name)
	assert.Equal(t, float32(10), product.Price)
	assert.Equal(t, "Category1", product.Category)
}

func TestCreateProduct(t *testing.T) {
	setupDatabaseProducts()
	e := echo.New()
	productData := `{
		"name": "NewProduct",
		"price": 50.0,
		"category_id": 1
	}`
	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(productData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := CreateProduct(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var response ProductResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "NewProduct", response.Name)
	assert.Equal(t, float32(50), response.Price)
	assert.Equal(t, "Category1", response.Category)

	var product models.Product
	database.DB.First(&product, "name = ?", "NewProduct")
	assert.Equal(t, "NewProduct", product.Name)
	assert.Equal(t, float32(50), product.Price)
	assert.Equal(t, uint(1), product.CategoryID)
}

func TestDeleteProduct(t *testing.T) {
	setupDatabaseProducts()

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues("1")

	err := DeleteProduct(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)

	var product models.Product
	result := database.DB.First(&product, "id=?", 1)
	assert.Error(t, result.Error)
}

func TestUpdateProduct(t *testing.T) {
	setupDatabaseProducts()
	e := echo.New()
	productData := `{
		"name": "UpdatedProduct",
		"price": 150.0,
		"category_id": 2
	}`
	req := httptest.NewRequest(http.MethodPut, "/products", strings.NewReader(productData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues("1")

	err := UpdateProduct(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response ProductResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "UpdatedProduct", response.Name)
	assert.Equal(t, float32(150), response.Price)
	assert.Equal(t, "2", response.Category)

	var product models.Product
	database.DB.First(&product, "id = ?", 1)
	assert.Equal(t, "UpdatedProduct", product.Name)
	assert.Equal(t, float32(150), product.Price)
	assert.Equal(t, uint(2), product.CategoryID)
}
