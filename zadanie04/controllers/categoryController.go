package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"zadanie04/database"
	"zadanie04/models"
	"zadanie04/scopes"
	"zadanie04/utils"
)

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func CreateCategory(c echo.Context) error {
	category := new(models.Category)
	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if category.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Category name cannot be empty"})
	}

	// Sprawdź, czy kategoria już istnieje
	var existingCategory models.Category
	result := database.DB.Scopes(scopes.Name(category.Name)).First(&existingCategory)
	if result.Error == nil {
		// Kategoria już istnieje
		return c.JSON(http.StatusConflict, map[string]string{"error": "Category already exists"})
	}

	result = database.DB.Create(category)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}

	response := CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}

	return c.JSON(http.StatusCreated, response)
}

func GetCategories(c echo.Context) error {
	var categories []models.Category
	result := database.DB.Find(&categories)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}

	var categoryResponse []CategoryResponse
	for _, category := range categories {
		categoryResponse = append(categoryResponse, CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return c.JSON(http.StatusOK, categoryResponse)
}

func GetCategory(c echo.Context) error {
	id, err := utils.ValidateID(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	var category models.Category
	result := database.DB.First(&category, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Category not found"})
	}

	response := CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateCategory(c echo.Context) error {
	id, err := utils.ValidateID(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	category := new(models.Category)
	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if category.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Category name cannot be empty"})
	}

	var existingCategory models.Category
	result := database.DB.First(&existingCategory, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Category not found"})
	}

	existingCategory.Name = category.Name

	database.DB.Save(&existingCategory)

	response := CategoryResponse{
		ID:   existingCategory.ID,
		Name: existingCategory.Name,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteCategory(c echo.Context) error {
	id, err := utils.ValidateID(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	result := database.DB.Delete(&models.Category{}, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}

	return c.JSON(http.StatusNoContent, nil)
}
