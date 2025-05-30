package order

import (
	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

// InitRoutes ...
func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(corsMiddleware())

	order := r.Group("/orders")
	{
		order.GET("", h.GetAllOrders)
		order.GET("/active", h.GetActiveOrders)
		order.GET("/:id/status", h.GetOrderStatus)
		order.GET("/view/:id", h.GetOrderByPublicID)
	}

	r.POST("/checkout", h.Checkout)

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
