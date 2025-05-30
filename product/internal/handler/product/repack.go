package product

import (
	"encoding/json"
	"github.com/Deevins/lampshop-backend/product/internal/model"
	"github.com/Deevins/lampshop-backend/product/internal/service/sql"
	"github.com/samber/lo"
)

func toProductResponse(p *sql.Product) model.ProductFull {
	var attrs map[string]interface{}
	if err := json.Unmarshal(p.Attributes, &attrs); err != nil {
		attrs = make(map[string]interface{}) // fallback, чтобы не крашилось
	}

	return model.ProductFull{
		ID:          p.ID,
		SKU:         p.Sku,
		Name:        p.Name,
		Description: lo.FromPtr(p.Description),
		CategoryID:  p.CategoryID,
		IsActive:    p.IsActive,
		ImageURL:    lo.FromPtr(p.ImageUrl),
		Price:       p.Price.InexactFloat64(),
		StockQty:    int(p.StockQty),
		Attributes:  attrs,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func toCategoryResponse(c sql.Category) model.CategoryResponse {
	return model.CategoryResponse{
		ID:   c.ID,
		Name: c.Name,
	}
}
