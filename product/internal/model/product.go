package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
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
