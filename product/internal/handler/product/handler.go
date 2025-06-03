package product

import (
	"context"
	"github.com/Deevins/lampshop-backend/product/internal/model"
	"github.com/Deevins/lampshop-backend/product/internal/service/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Service interface {
	ListProducts(ctx context.Context) ([]*sql.Product, error)
	GetProduct(ctx context.Context, id uuid.UUID) (*sql.Product, error)
	CreateProduct(ctx context.Context, params *model.CreateProductRequest) error
	UpdateProduct(ctx context.Context, params *model.UpdateProductRequest) error
	DeleteProduct(ctx context.Context, id uuid.UUID) error

	ListCategories(ctx context.Context) ([]*sql.Category, error)
	CreateCategory(ctx context.Context, c *model.CategoryCreateRequest) error
	DeleteCategory(ctx context.Context, id uuid.UUID) error
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
	r.Use(corsMiddleware())

	products := r.Group("/products")
	{
		products.GET("", h.ListProducts)
		products.POST("", h.CreateProduct)
		products.GET("/:id", h.GetProductByID)
		products.PUT("/:id", h.UpdateProduct)
		products.DELETE("/:id", h.DeleteProduct)

		products.GET("/upload-url", h.GetUploadURL)
		products.POST("/notify-upload", h.NotifyUpload)
	}

	categories := r.Group("/categories")
	{
		categories.POST("", h.CreateCategory)
		categories.GET("", h.ListCategories)
		categories.DELETE("/:id", h.DeleteCategory)
	}

	return r
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		hdr := c.Writer.Header()

		if _, ok := hdr["Access-Control-Allow-Origin"]; !ok {
			hdr.Set("Access-Control-Allow-Origin", "*")
		}
		if _, ok := hdr["Access-Control-Allow-Methods"]; !ok {
			hdr.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		}
		if _, ok := hdr["Access-Control-Allow-Headers"]; !ok {
			hdr.Set("Access-Control-Allow-Headers", "*")
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
