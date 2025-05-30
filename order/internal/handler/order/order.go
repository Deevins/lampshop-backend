package order

import (
	"github.com/Deevins/lampshop-backend/order/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (h *Handler) Register(r *gin.RouterGroup) {
	orders := r.Group("/orders")
	{
		orders.GET("", h.GetAllOrders)
		orders.GET("/active", h.GetActiveOrders)
		orders.GET("/:id/status", h.GetOrderStatus)
	}
	r.POST("/checkout", h.CreateOrder)
}

// --- Handlers ---

func (h *Handler) GetAllOrders(c *gin.Context) {
	orders, err := h.service.GetAllOrders(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *Handler) GetActiveOrders(c *gin.Context) {
	orders, err := h.service.GetActiveOrders(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch active orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *Handler) GetOrderStatus(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	status, err := h.service.GetOrderStatus(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	c.JSON(http.StatusOK, model.OrderStatusResponse{
		OrderID: status.OrderID,
		Status:  status.Status,
	})
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var req model.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if len(req.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order must contain items"})
		return
	}

	err := h.service.CreateOrder(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "order created"})
}
