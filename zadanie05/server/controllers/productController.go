package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"productsServer/database"
	"productsServer/models"
	"strings"
)

type ProductResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}

func GetProducts(c echo.Context) error {
	var products []models.Product

	result := database.DB.Find(&products)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}

	var productResponse []ProductResponse
	for _, product := range products {
		productResponse = append(productResponse, ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
		})
	}

	return c.JSON(http.StatusOK, productResponse)
}

func GetProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	result := database.DB.First(&product, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}

	response := ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
	}

	return c.JSON(http.StatusOK, response)
}

func CreateProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	fmt.Println(product)

	if product.Price == 0.0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Price must be greater than 0"})
	}

	if product.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Name cannot be null"})
	}

	product.Name = strings.TrimSpace(product.Name)

	var existingProduct models.Product
	result := database.DB.Where("name = ?", product.Name).First(&existingProduct)
	if result.Error == nil {
		// Kategoria już istnieje
		return c.JSON(http.StatusConflict, map[string]string{"error": "Product already exists"})
	}

	database.DB.Create(&product)
	return c.JSON(http.StatusCreated, product)
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	result := database.DB.Delete(&models.Product{}, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}

	return c.JSON(http.StatusNoContent, nil)
}
