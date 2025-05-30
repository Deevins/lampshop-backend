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
	c.JSON(http.StatusOK, gin.H{"message": "Get product status"})
}

func (h *Handler) GetOrderByPublicID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get product by public ID"})
}

func (h *Handler) Checkout(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Checkout"}) }
