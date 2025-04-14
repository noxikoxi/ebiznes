package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"zadanie04/database"
	"zadanie04/models"
	"zadanie04/scopes"
)

type CartResponse struct {
	ID        uint                `json:"id"`
	CartItems []CartItemsResponse `json:"cart_items"`
	Total     float32             `json:"total"`
}

type CartItemsResponse struct {
	Product   string  `json:"product"`
	Quantity  uint    `json:"quantity"`
	TotalCost float32 `json:"total_cost"`
}

type CartItemResponse struct {
	CartID    uint `json:"cart_id"`
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

func GetCarts(c echo.Context) error {
	var carts []models.Cart
	result := database.DB.Preload("CartItems.Product").Find(&carts)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}

	var totalCost float32

	var cartResponse []CartResponse

	for _, cart := range carts {
		totalCost = 0.0
		var cartItemsResponse []CartItemsResponse

		for _, item := range cart.CartItems {
			itemCost := float32(item.Quantity) * item.Product.Price
			totalCost += itemCost
			cartItemsResponse = append(cartItemsResponse, CartItemsResponse{
				Product:   item.Product.Name,
				Quantity:  item.Quantity,
				TotalCost: itemCost,
			})
		}

		cartResponse = append(cartResponse, CartResponse{
			ID:        cart.ID,
			CartItems: cartItemsResponse,
			Total:     totalCost,
		})
	}

	return c.JSON(http.StatusOK, cartResponse)
}

func CreateCart(c echo.Context) error {
	var cart models.Cart
	// Pusty koszyk nie potrzebuje żadnych danych
	result := database.DB.Create(&cart)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}

	response := CartResponse{
		ID:        cart.ID,
		CartItems: []CartItemsResponse{},
	}

	return c.JSON(http.StatusCreated, response)
}

func GetCart(c echo.Context) error {
	id := c.Param("id")
	cart := models.Cart{}

	result := database.DB.Preload("CartItems.Product").First(&cart, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Cart not found"})
	}

	var totalCost float32
	var cartResponse CartResponse
	var cartItemsResponse []CartItemsResponse
	for _, item := range cart.CartItems {
		itemCost := float32(item.Quantity) * item.Product.Price
		totalCost += itemCost
		cartItemsResponse = append(cartItemsResponse, CartItemsResponse{
			Product:   item.Product.Name,
			Quantity:  item.Quantity,
			TotalCost: float32(item.Quantity) * item.Product.Price,
		})
	}

	cartResponse.ID = cart.ID
	cartResponse.CartItems = cartItemsResponse
	cartResponse.Total = totalCost

	return c.JSON(http.StatusOK, cartResponse)
}

func DeleteCart(c echo.Context) error {
	id := c.Param("id")

	result := database.DB.Delete(&models.Cart{}, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}

	return c.JSON(http.StatusNoContent, nil)
}

func AddCartItem(c echo.Context) error {
	cartID := c.Param("cart_id")
	productID := c.Param("product_id")
	quantity := c.Param("quantity")

	parsedCartID, err1 := strconv.ParseUint(cartID, 10, 32)
	parsedProductID, err2 := strconv.ParseUint(productID, 10, 32)
	parsedQuantity, err3 := strconv.ParseUint(quantity, 10, 32)

	if err1 != nil || err2 != nil || err3 != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "cart_id, product_id and quantity have to be uint type"})
	}

	cartItem := models.CartItem{
		CartID:    uint(parsedCartID),
		ProductID: uint(parsedProductID),
		Quantity:  uint(parsedQuantity),
	}

	var cart models.Cart

	result := database.DB.Preload("CartItems.Product").First(&cart, cartID)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Cart not found"})
	}

	for _, item := range cart.CartItems {
		if item.Product.ID == cartItem.ProductID {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Product already exists"})
		}
	}

	cart.CartItems = append(cart.CartItems, cartItem)

	response := CartItemResponse{}
	response.CartID = cart.ID
	response.Quantity = cartItem.Quantity
	response.ProductID = cartItem.ProductID

	database.DB.Save(&cart)

	return c.JSON(http.StatusOK, response)
}

func UpdateCartItem(c echo.Context) error {
	cartID := c.Param("cart_id")
	productID := c.Param("product_id")
	quantity := c.Param("quantity")

	parsedCartID, err1 := strconv.ParseUint(cartID, 10, 32)
	parsedProductID, err2 := strconv.ParseUint(productID, 10, 32)
	parsedQuantity, err3 := strconv.ParseUint(quantity, 10, 32)

	if err1 != nil || err2 != nil || err3 != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "cart_id, product_id and quantity have to be uint type"})
	}

	cart := models.Cart{}
	var cartItem models.CartItem

	result := database.DB.Preload("CartItems").First(&cart, cartID)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Cart not found"})
	}

	productFound := false

	for _, item := range cart.CartItems {
		if item.ProductID == uint(parsedProductID) {
			productFound = true
			break
		}
	}

	if !productFound {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Product not found"})
	}

	result = database.DB.Scopes(scopes.CartID(uint(parsedCartID)), scopes.CartProductID(uint(parsedProductID))).First(&cartItem)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Cart Item not found"})
	}

	cartItem.Quantity = uint(parsedQuantity)

	database.DB.Save(&cartItem)

	response := CartItemResponse{}
	response.CartID = cart.ID
	response.Quantity = cartItem.Quantity
	response.ProductID = cartItem.ProductID

	return c.JSON(http.StatusOK, response)
}

func DeleteCartItem(c echo.Context) error {
	cartID := c.Param("cart_id")
	productID := c.Param("product_id")

	parsedProductID, err := strconv.ParseUint(productID, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "product_id has to be uint type"})
	}

	cart := models.Cart{}
	result := database.DB.Preload("CartItems").First(&cart, cartID)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Cart not found"})
	}

	var cartItem models.CartItem

	for _, item := range cart.CartItems {
		if item.ProductID == uint(parsedProductID) {
			cartItem = item
			break
		}
	}

	if cartItem.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found in cart"})
	}

	result = database.DB.Delete(&models.CartItem{}, cartItem.ID)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}

	return c.JSON(http.StatusNoContent, nil)
}
