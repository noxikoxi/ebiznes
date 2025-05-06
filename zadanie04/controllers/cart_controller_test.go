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
	"testing"
	"zadanie04/database"
	"zadanie04/models"
)

func setupDatabaseCarts() {
	// Tworzenie bazy w pamięci
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("could not connect to in-memory database")
	}

	// Migracje, żeby tworzyć tabele
	db.AutoMigrate(&models.Product{}, &models.Cart{}, &models.CartItem{}, &models.Category{})

	db.Create(&models.Category{ID: 1, Name: "Category1"})
	db.Create(&models.Category{ID: 2, Name: "Category1"})
	db.Create(&models.Category{ID: 3, Name: "Category1"})

	db.Create(&models.Product{ID: 1, Name: "Item1", Price: 10, CategoryID: 1})
	db.Create(&models.Product{ID: 2, Name: "Item2", Price: 100, CategoryID: 1})
	db.Create(&models.Product{ID: 3, Name: "Item3", Price: 1000, CategoryID: 2})
	db.Create(&models.Product{ID: 4, Name: "Item4", Price: 10, CategoryID: 2})
	db.Create(&models.Product{ID: 5, Name: "Item5", Price: 10, CategoryID: 3})
	db.Create(&models.Product{ID: 6, Name: "Item6", Price: 1000, CategoryID: 3})

	db.Create(&models.Cart{ID: 1})
	db.Create(&models.Cart{ID: 2})
	db.Create(&models.Cart{ID: 3})

	db.Create(&models.CartItem{ID: 1, CartID: 1, ProductID: 1, Quantity: 1})
	db.Create(&models.CartItem{ID: 2, CartID: 1, ProductID: 2, Quantity: 4})
	db.Create(&models.CartItem{ID: 3, CartID: 1, ProductID: 3, Quantity: 5})
	db.Create(&models.CartItem{ID: 4, CartID: 2, ProductID: 1, Quantity: 2})
	db.Create(&models.CartItem{ID: 5, CartID: 2, ProductID: 5, Quantity: 2})
	db.Create(&models.CartItem{ID: 6, CartID: 3, ProductID: 6, Quantity: 99})

	database.DB = db
}

func TestGetCarts(t *testing.T) {
	setupDatabaseCarts()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/carts", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := GetCarts(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	var carts []CartResponse
	err = json.Unmarshal(rec.Body.Bytes(), &carts)
	assert.NoError(t, err)
	expectedCount := 3
	assert.Equal(t, expectedCount, len(carts))
}

func TestGetCart(t *testing.T) {
	setupDatabaseCarts()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/carts/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	err := GetCart(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	var cart CartResponse
	err = json.Unmarshal(rec.Body.Bytes(), &cart)
	assert.NoError(t, err)

	assert.Equal(t, uint(1), cart.ID)
	assert.Equal(t, float32(5410), cart.Total)
	expectedItems := 3
	assert.Equal(t, expectedItems, len(cart.CartItems))
	assert.Equal(t, uint(1), cart.CartItems[0].Quantity)
	assert.Equal(t, "Item1", cart.CartItems[0].Product)
	assert.Equal(t, float32(10), cart.CartItems[0].TotalCost)
}

func TestCreateCart(t *testing.T) {
	setupDatabaseCarts()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/carts", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := CreateCart(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	var response CartResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, uint(4), response.ID)
	assert.Empty(t, response.CartItems)
}

func TestDeleteCart(t *testing.T) {
	setupDatabaseCarts()
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/carts", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues("1")

	err := DeleteCart(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)
	var cart models.Cart
	result := database.DB.First(&cart, "id=?", 1)
	assert.Error(t, result.Error)
}

func TestAddProduct(t *testing.T) {
	setupDatabaseCarts()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/carts/3/5/50", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("cart_id", "product_id", "quantity")
	c.SetParamValues("3", "5", "50")

	err := AddCartItem(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	var response CartItemResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, uint(3), response.CartID)
	assert.Equal(t, uint(5), response.ProductID)
	assert.Equal(t, uint(50), response.Quantity)

	var cartItem models.CartItem
	err = json.Unmarshal(rec.Body.Bytes(), &cartItem)
	result := database.DB.First(&cartItem, "id=?", 7)
	assert.NoError(t, result.Error)
	assert.Equal(t, uint(3), cartItem.CartID)
}

func TestUpdateCartItem(t *testing.T) {
	setupDatabaseCarts()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/carts/1/1/25", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("cart_id", "product_id", "quantity")
	c.SetParamValues("1", "1", "25")

	err := UpdateCartItem(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	var response CartItemResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), response.CartID)
	assert.Equal(t, uint(1), response.ProductID)
	assert.Equal(t, uint(25), response.Quantity)

	var cartItem models.CartItem
	err = json.Unmarshal(rec.Body.Bytes(), &cartItem)
	result := database.DB.First(&cartItem, "id=?", 1)
	assert.NoError(t, result.Error)
	assert.Equal(t, uint(25), cartItem.Quantity)
}

func TestDeleteCartItem(t *testing.T) {
	setupDatabaseCarts()
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/carts/1/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("cart_id", "product_id")
	c.SetParamValues("1", "1")

	err := DeleteCartItem(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)

	var cartItem models.CartItem
	err = json.Unmarshal(rec.Body.Bytes(), &cartItem)
	result := database.DB.First(&cartItem, "id=?", 1)
	assert.Error(t, result.Error)
}
