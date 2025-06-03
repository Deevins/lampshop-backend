package order

import (
	"fmt"
	"github.com/Deevins/lampshop-backend/order/internal/model"
	"github.com/Deevins/lampshop-backend/order/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (h *Handler) GetAllOrders(c *gin.Context) {
	orders, err := h.service.GetAllOrders(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch orders", "errorDetail": err.Error()})
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
		Amount:  status.Amount,
		Status:  status.Status,
	})
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var req model.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log.Errorw("failed to bind create order request", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	fmt.Printf("%+v", req)

	if len(req.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order must contain items"})
		return
	}

	orderID, err := h.service.CreateOrder(c.Request.Context(), req)
	if err != nil {
		logger.Log.Errorw("failed to create order", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"order_id": orderID, "message": "order created"})
}

func (h *Handler) UpdateOrderStatus(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	var req *model.UpdateOrderStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	req.OrderID = id

	err = h.service.UpdateStatus(c.Request.Context(), req.OrderID, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "order status updated"})
}
