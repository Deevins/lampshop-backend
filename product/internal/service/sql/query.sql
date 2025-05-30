-- name: GetProductByID :one
SELECT * FROM products WHERE id = $1;

-- name: ListProducts :many
SELECT * FROM products ORDER BY created_at DESC;

-- name: CreateProduct :exec
INSERT INTO products (
    id, sku, name, description, category_id, is_active, image_url,
    price, stock_qty, attributes, created_at, updated_at
) VALUES (
             $1, $2, $3, $4, $5, $6, $7,
             $8, $9, $10, $11, $12
         );

-- name: UpdateProduct :exec
UPDATE products SET
                    name = $2,
                    description = $3,
                    category_id = $4,
                    is_active = $5,
                    image_url = $6,
                    price = $7,
                    stock_qty = $8,
                    attributes = $9,
                    updated_at = $10
WHERE id = $1;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;


-- name: ListCategories :many
SELECT id, name FROM categories ORDER BY name;

-- name: CreateCategory :exec
INSERT INTO categories (id, name) VALUES ($1, $2);

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1;