package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllOrders(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all orders"})
}

func (h *Handler) GetActiveOrders(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get active orders"})
}

func (h *Handler) GetOrderStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get order status"})
}

func (h *Handler) GetOrderByPublicID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get order by public ID"})
}

func (h *Handler) GetCart(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Get cart"}) }

func (h *Handler) AddItemToCart(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Add item"})
}

func (h *Handler) RemoveItemFromCart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Remove item"})
}

func (h *Handler) Checkout(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Checkout"}) }
