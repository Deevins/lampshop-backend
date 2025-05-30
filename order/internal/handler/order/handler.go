package order

import (
	"context"
	"github.com/Deevins/lampshop-backend/order/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Service interface {
	GetAllOrders(ctx context.Context) ([]model.Order, error)
	GetActiveOrders(ctx context.Context) ([]model.Order, error)
	GetOrderStatus(ctx context.Context, id uuid.UUID) (model.OrderStatusResponse, error)
	CreateOrder(ctx context.Context, req model.CreateOrderRequest) error
}

// Handler ...
type Handler struct {
	service Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{
		service: svc,
	}
}

// InitRoutes ...
func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	orders := r.Group("/orders")
	{
		orders.GET("", h.GetAllOrders)
		orders.GET("/active", h.GetActiveOrders)
		orders.GET("/:id/status", h.GetOrderStatus)
	}
	r.POST("/checkout", h.CreateOrder)

	r.Use(corsMiddleware())

	return r
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
