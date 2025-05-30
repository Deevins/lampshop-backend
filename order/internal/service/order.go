package service

import (
	"context"
	"github.com/Deevins/lampshop-backend/order/internal/model"
	"github.com/Deevins/lampshop-backend/order/internal/service/sql"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

type OrderService struct {
	db *pgxpool.Pool
}

func NewOrderService(db *pgxpool.Pool) *OrderService {
	return &OrderService{db: db}
}

func (s *OrderService) CreateOrder(ctx context.Context, req model.CreateOrderRequest) error {
	repo := sql.New(s.db)

	total := lo.Reduce(req.Items, func(acc float64, item model.OrderItemInput, _ int) float64 {
		return acc + float64(item.Qty)*item.UnitPrice
	}, 0.0)

	if total != req.Payment.Amount {
		return errors.New("products total should be equal to payment amount")
	}

	return repo.CreateOrder(ctx, &sql.CreateOrderParams{
		ID:       uuid.New(),
		Status:   sql.PaymentStatusPending,
		Total:    decimal.NewFromFloat(total),
		IsActive: true,
	})
}

func (s *OrderService) GetAllOrders(ctx context.Context) ([]model.Order, error) {
	repo := sql.New(s.db)

	resp, err := repo.GetAllOrders(ctx)
	if err != nil {
		return nil, err
	}

	return toModelOrderList(resp), nil
}

func (s *OrderService) GetActiveOrders(ctx context.Context) ([]model.Order, error) {
	repo := sql.New(s.db)

	resp, err := repo.GetActiveOrders(ctx)
	if err != nil {
		return nil, err
	}

	return toModelOrderList(resp), nil
}

func (s *OrderService) GetOrderStatus(ctx context.Context, id uuid.UUID) (model.OrderStatusResponse, error) {
	repo := sql.New(s.db)

	resp, err := repo.GetOrderStatus(ctx, id)
	if err != nil {
		return model.OrderStatusResponse{}, err
	}

	return model.OrderStatusResponse{
		OrderID: resp.ID,
		Status:  string(resp.Status),
	}, nil
}
