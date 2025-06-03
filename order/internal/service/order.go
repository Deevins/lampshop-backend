package service

import (
	"context"
	"fmt"
	"github.com/Deevins/lampshop-backend/order/internal/model"
	"github.com/Deevins/lampshop-backend/order/internal/service/sql"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

func (s *OrderService) CreateOrder(ctx context.Context, req model.CreateOrderRequest) (uuid.UUID, error) {
	repo := sql.New(s.db)

	total := lo.Reduce(req.Items, func(acc float64, item model.OrderItemInput, _ int) float64 {
		return acc + float64(item.Qty)*item.UnitPrice
	}, 0.0)

	orderID, err := repo.CreateOrder(ctx, &sql.CreateOrderParams{
		ID:                uuid.New(),
		Status:            sql.PaymentStatusPending,
		Total:             decimal.NewFromFloat(total),
		IsActive:          true,
		CustomerFirstName: req.Customer.FirstName,
		CustomerLastName:  req.Customer.LastName,
		CustomerEmail:     req.Customer.Email,
		CustomerPhone:     req.Customer.Phone,
		Address:           req.Customer.Address,
	})
	if err != nil {
		return uuid.Nil, err
	}

	for _, item := range req.Items {
		err = repo.AddOrderItem(ctx, &sql.AddOrderItemParams{
			ID:        uuid.New(),
			OrderID:   orderID,
			ProductID: item.ProductID,
			Qty:       int32(item.Qty),
			UnitPrice: decimal.NewFromFloat(item.UnitPrice),
		})
	}

	return orderID, nil
}

func (s *OrderService) GetAllOrders(ctx context.Context) ([]model.Order, error) {
	repo := sql.New(s.db)

	ordersRows, err := repo.GetAllOrders(ctx)
	if err != nil {
		return nil, err
	}

	if len(ordersRows) == 0 {
		return []model.Order{}, nil
	}

	var orderIDs []uuid.UUID
	for _, o := range ordersRows {
		orderIDs = append(orderIDs, o.ID)
	}

	items, err := repo.GetOrderItemsByOrderIDs(ctx, orderIDs)
	if err != nil {
		return nil, fmt.Errorf("GetOrderItemsByOrderIDs failed: %w", err)
	}

	itemsMap := make(map[string][]model.OrderItem, len(orderIDs))
	for _, ir := range items {
		var priceFloat float64
		switch v := any(ir.OrderItemPrice).(type) {
		case decimal.Decimal:
			pf, _ := v.Float64()
			priceFloat = pf
		default:
			priceFloat = 0
		}

		oi := model.OrderItem{
			ID:        ir.OrderItemID,
			OrderID:   ir.OrderItemOrderID,
			ProductID: ir.OrderItemProductID,
			Quantity:  float64(ir.OrderItemQuantity),
			Price:     priceFloat,
		}
		key := ir.OrderItemOrderID.String()
		itemsMap[key] = append(itemsMap[key], oi)
	}

	var result []model.Order
	for _, or := range ordersRows {
		var totalFloat float64
		switch v := any(or.Total).(type) {
		case decimal.Decimal:
			tf, _ := v.Float64()
			totalFloat = tf
		default:
			totalFloat = 0
		}

		fullName := or.CustomerFirstName + " " + or.CustomerLastName

		o := model.Order{
			ID:        or.ID,
			Status:    string(or.Status),
			FullName:  fullName,
			Items:     itemsMap[or.ID.String()],
			Total:     totalFloat,
			IsActive:  or.IsActive,
			CreatedAt: or.CreatedAt,
			UpdatedAt: or.UpdatedAt,
		}
		result = append(result, o)
	}

	return result, nil
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
		if errors.Is(err, pgx.ErrNoRows) {
			return model.OrderStatusResponse{}, fmt.Errorf("order not found")
		}

		return model.OrderStatusResponse{}, err
	}

	return model.OrderStatusResponse{
		OrderID: resp.ID,
		Amount:  resp.Total.InexactFloat64(),
		Status:  string(resp.Status),
	}, nil
}

func (s *OrderService) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	repo := sql.New(s.db)

	if err := repo.UpdateOrderStatus(ctx, &sql.UpdateOrderStatusParams{
		Status: sql.PaymentStatus(status),
		ID:     id,
	}); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("order not found")
		}

		return err
	}

	return nil
}
