package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"zadanie04/database"
	"zadanie04/models"
	"zadanie04/scopes"
)

type ProductResponse struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Category string  `json:"category"`
}

func GetProducts(c echo.Context) error {
	var products []models.Product
	minPriceStr := c.QueryParam("minPrice")     // np. minPrice=100
	categoryIDStr := c.QueryParam("categoryID") // np. categoryID=1

	filters := scopes.GetCategoryPriceFilters(minPriceStr, categoryIDStr)

	result := database.DB.Preload("Category").Scopes(filters...).Find(&products)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}

	var productResponse []ProductResponse
	for _, product := range products {
		productResponse = append(productResponse, ProductResponse{
			ID:       product.ID,
			Name:     product.Name,
			Price:    product.Price,
			Category: product.Category.Name,
		})
	}

	return c.JSON(http.StatusOK, productResponse)
}

func GetProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	result := database.DB.Preload("Category").First(&product, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}

	response := ProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Category: product.Category.Name,
	}

	return c.JSON(http.StatusOK, response)
}

func CreateProduct(c echo.Context) error {
	product := new(models.Product)

	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Walidacja
	if product.CategoryID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Category_id cannot be null"})
	}

	if product.Price == 0.0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Price must be greater than 0"})
	}

	var existingProduct models.Product
	result := database.DB.Scopes(scopes.Name(product.Name)).First(&existingProduct)
	if result.Error == nil {
		// Kategoria już istnieje
		return c.JSON(http.StatusConflict, map[string]string{"error": "Product already exists"})
	}

	result = database.DB.Create(product)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}

	// Ładowanie powiązanej kategorii
	result = database.DB.Preload("Category").First(&product, product.ID)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error loading product's category"})
	}

	response := ProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Category: product.Category.Name,
	}

	return c.JSON(http.StatusCreated, response)
}

func UpdateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if product.Name == "" || product.Price == 0 || product.CategoryID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Product name cannot be empty, price and categoryID cannot be 0"})
	}

	var existingProduct models.Product
	result := database.DB.Preload("Category").First(&existingProduct, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}

	existingProduct.Name = product.Name
	existingProduct.Price = product.Price
	existingProduct.CategoryID = product.CategoryID
	existingProduct.Category = models.Category{}

	if err := database.DB.Save(&existingProduct).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := ProductResponse{
		ID:       existingProduct.ID,
		Name:     existingProduct.Name,
		Price:    existingProduct.Price,
		Category: strconv.Itoa(int(existingProduct.CategoryID)),
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")

	productID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	result := database.DB.Delete(&models.Product{}, productID)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}

	return c.JSON(http.StatusNoContent, nil)
}
