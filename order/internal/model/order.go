package model

import (
	"github.com/google/uuid"
	"time"
)

type PaymentProvider string

const (
	Stripe PaymentProvider = "stripe"
	Sbp    PaymentProvider = "sbp"
)

type Order struct {
	ID        uuid.UUID `json:"id"`
	Status    string    `json:"status"`
	Total     float64   `json:"total"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Customer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

type CreateOrderRequest struct {
	Items    []OrderItemInput `json:"items"`
	Customer Customer         `json:"customer"`
	Payment  PaymentInput     `json:"payment"`
}

type OrderItemInput struct {
	ProductID uuid.UUID `json:"product_id"`
	Qty       int       `json:"qty"`
	UnitPrice float64   `json:"unit_price"`
}

type PaymentInput struct {
	Provider       PaymentProvider `json:"provider"`
	Amount         float64         `json:"amount"`
	TransactionRef string          `json:"transaction_ref"`
}

type OrderStatusResponse struct {
	OrderID uuid.UUID `json:"order_id"`
	Amount  float64   `json:"amount"`
	Status  string    `json:"status"`
}
