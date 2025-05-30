package product

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

	products := r.Group("/products")
	{
		products.GET("", h.ListProducts)
		products.POST("", h.CreateProduct)
		products.GET("/:id", h.GetProductByID)
		products.PATCH("/:id", h.UpdateProduct)
		products.DELETE("/:id", h.DeleteProduct)

		products.GET("/upload-url", h.GetUploadURL)
		products.POST("/notify-upload", h.NotifyUpload)
	}

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
