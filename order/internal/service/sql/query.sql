-- name: GetAllOrders :many
SELECT *
FROM orders
ORDER BY created_at DESC;

-- name: GetActiveOrders :many
SELECT *
FROM orders
WHERE is_active = true
ORDER BY created_at DESC;

-- name: GetOrderStatus :one
SELECT id, status
FROM orders
WHERE id = $1;

-- name: CreateOrder :one
INSERT INTO orders (id, status, total, is_active)
VALUES ($1, $2, $3, $4) RETURNING id;


-- name: AddOrderItem :exec
INSERT INTO order_items (id, order_id, product_id, qty, unit_price)
VALUES ($1, $2, $3, $4, $5);


-- name: CreatePayment :exec
INSERT INTO payments (id, order_id, provider, status, amount, transaction_ref)
VALUES ($1, $2, $3, $4, $5, $6);
