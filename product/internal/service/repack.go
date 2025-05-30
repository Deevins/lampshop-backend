package service

import (
	"encoding/json"
	"github.com/Deevins/lampshop-backend/product/internal/model"
	"github.com/Deevins/lampshop-backend/product/internal/service/sql"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"time"
)

func mapUpdateRequestToSQL(id uuid.UUID, req *model.UpdateProductRequest) *sql.UpdateProductParams {
	var result = &sql.UpdateProductParams{}
	result.ID = id
	result.UpdatedAt = time.Now()

	result.Name = req.Name
	result.Description = lo.ToPtr(req.Description)
	result.CategoryID = req.CategoryID
	result.IsActive = req.IsActive
	result.ImageUrl = lo.ToPtr(req.ImageURL)
	result.StockQty = int32(req.StockQty)

	result.Price = decimal.NewFromFloat(req.Price)

	raw, err := json.Marshal(req.Attributes)
	if err != nil {
		result.Attributes = nil
	}

	result.Attributes = raw

	return result
}
