package model

import (
	"time"

	"github.com/google/uuid"
)

type ProductFull struct {
	ID          uuid.UUID              `json:"id"`
	SKU         string                 `json:"sku"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	CategoryID  uuid.UUID              `json:"category_id"`
	IsActive    bool                   `json:"is_active"`
	ImageURL    string                 `json:"image_url"`
	Price       float64                `json:"price"`
	StockQty    int                    `json:"stock_qty"`
	Attributes  map[string]interface{} `json:"attributes"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}

type Category struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CreateProductRequest struct {
	SKU         string                 `json:"sku" binding:"required"`
	Name        string                 `json:"name" binding:"required"`
	Description string                 `json:"description"`
	CategoryID  uuid.UUID              `json:"category_id" binding:"required"`
	IsActive    bool                   `json:"is_active"`
	ImageURL    string                 `json:"image_url"`
	Price       float64                `json:"price" binding:"required"`
	StockQty    int                    `json:"stock_qty" binding:"required"`
	Attributes  map[string]interface{} `json:"attributes" binding:"required"`
}

type UpdateProductRequest struct {
	ID          uuid.UUID              `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	CategoryID  uuid.UUID              `json:"category_id"`
	IsActive    bool                   `json:"is_active"`
	ImageURL    string                 `json:"image_url"`
	Price       float64                `json:"price"`
	StockQty    int                    `json:"stock_qty"`
	Attributes  map[string]interface{} `json:"attributes"`
}

type ProductResponse struct {
	ID          uuid.UUID              `json:"id"`
	SKU         string                 `json:"sku"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	CategoryID  uuid.UUID              `json:"category_id"`
	IsActive    bool                   `json:"is_active"`
	ImageURL    string                 `json:"image_url"`
	Price       float64                `json:"price"`
	StockQty    int                    `json:"stock_qty"`
	Attributes  map[string]interface{} `json:"attributes"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}

type CategoryCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

type CategoryResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
