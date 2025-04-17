package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"productsServer/database"
	"productsServer/models"
	"strconv"
	"time"
)

type CartItemResponse struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity uint    `json:"quantity"`
}

type PaymentResponse struct {
	Items []CartItemResponse `json:"items"`
	Total float32            `json:"total"`
	Date  time.Time          `json:"date"`
}

type PaymentRequest struct {
	Items []CartItemRequest `json:"items"`
	Total float32           `json:"total"`
}

type CartItemRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

func createPaymentResponse(id string) (PaymentResponse, error) {
	payment := models.Payment{}
	result := database.DB.Preload("Cart.CartItems.Product").Find(&payment, id)
	if result.Error != nil {
		return PaymentResponse{}, result.Error
	}

	var cartItems []CartItemResponse

	for _, item := range payment.Cart.CartItems {
		cartItems = append(cartItems, CartItemResponse{
			Name:     item.Product.Name,
			Price:    item.Product.Price,
			Quantity: item.Quantity,
		})
	}

	paymentResponse := PaymentResponse{
		Items: cartItems,
		Total: payment.Total,
		Date:  payment.CreatedAt,
	}

	return paymentResponse, nil
}

func GetPaymentWithCart(c echo.Context) error {
	id := c.Param("id")

	payment, err := createPaymentResponse(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Payment not found"})
	}

	return c.JSON(http.StatusOK, payment)
}

func GetPaymentsWithCarts(c echo.Context) error {
	var payments []models.Payment

	result := database.DB.Preload("Cart.CartItems.Product").Find(&payments)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}

	var PaymentResponses []PaymentResponse

	for _, payment := range payments {
		var cartItems []CartItemResponse
		for _, item := range payment.Cart.CartItems {
			cartItems = append(cartItems, CartItemResponse{
				Name:     item.Product.Name,
				Price:    item.Product.Price,
				Quantity: item.Quantity,
			})
		}
		PaymentResponses = append(PaymentResponses, PaymentResponse{
			Items: cartItems,
			Total: payment.Total,
			Date:  payment.CreatedAt,
		})
	}
	return c.JSON(http.StatusOK, PaymentResponses)
}

func CreatePaymentWithCart(c echo.Context) error {
	var cart models.Cart
	result := database.DB.Create(&cart)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}

	var request = new(PaymentRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if request.Total == 0.0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Total cannot be zero"})
	}

	if len(request.Items) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Items cannot be empty"})
	}

	for _, item := range request.Items {
		cartItem := models.CartItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			CartID:    cart.ID,
		}
		database.DB.Create(&cartItem)
	}

	payment := models.Payment{
		CartID: cart.ID,
		Total:  request.Total,
	}

	database.DB.Create(&payment)

	response, err := createPaymentResponse(strconv.Itoa(int(payment.ID)))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, response)
}
