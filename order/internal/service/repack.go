package service

import (
	"github.com/Deevins/lampshop-backend/order/internal/model"
	"github.com/Deevins/lampshop-backend/order/internal/service/sql"
)

func toModelOrderList(sqlOrders []*sql.Order) []model.Order {
	var result []model.Order
	for _, o := range sqlOrders {
		result = append(result, toModelOrder(o))
	}
	return result
}

func toModelOrder(o *sql.Order) model.Order {
	return model.Order{
		ID:        o.ID,
		Status:    string(o.Status),
		Total:     o.Total.InexactFloat64(),
		IsActive:  o.IsActive,
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}
}
