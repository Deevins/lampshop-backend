-- name: GetAllOrders :many
SELECT *
FROM orders
ORDER BY created_at DESC;

-- 2) Получить все элементы сразу для массива order_id
--    Мы передаём массив UUID, а SQL возвращает все rows из order_items,
--    у которых order_id = ANY($1).
-- name: GetOrderItemsByOrderIDs :many
SELECT id         AS order_item_id,
       order_id   AS order_item_order_id,
       product_id AS order_item_product_id,
       qty        AS order_item_quantity,
       unit_price AS order_item_price
FROM order_items
WHERE order_id = ANY ($1::uuid[]);

-- name: GetActiveOrders :many
SELECT *
FROM orders
WHERE is_active = true
ORDER BY created_at DESC;

-- name: GetOrderStatus :one
SELECT id, status, total
FROM orders
WHERE id = $1;

-- name: CreateOrder :one
INSERT INTO orders (id, status, total, is_active, customer_first_name, customer_last_name,
                    customer_email, customer_phone, address)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id;


-- name: AddOrderItem :exec
INSERT INTO order_items (id, order_id, product_id, qty, unit_price)
VALUES ($1, $2, $3, $4, $5);


-- name: CreatePayment :exec
INSERT INTO payments (id, order_id, provider, status, amount, transaction_ref)
VALUES ($1, $2, $3, $4, $5, $6);


-- name: UpdateOrderStatus :exec
UPDATE orders
SET status = $1
where id = $2;