-- name: GetProductByID :one
SELECT *
FROM products
WHERE id = $1;

-- name: ListProducts :many
SELECT *
FROM products
ORDER BY created_at DESC;

-- name: CreateProduct :exec
INSERT INTO products (id, sku, name, description, category_id, is_active, image_url,
                      price, stock_qty, attributes, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7,
        $8, $9, $10, $11, $12);

-- name: UpdateProduct :exec
UPDATE products
SET name        = CASE WHEN name IS DISTINCT FROM $2 THEN $2 ELSE name END,
    description = CASE WHEN description IS DISTINCT FROM $3 THEN $3 ELSE description END,
    category_id = CASE WHEN category_id IS DISTINCT FROM $4 THEN $4 ELSE category_id END,
    is_active   = CASE WHEN is_active IS DISTINCT FROM $5 THEN $5 ELSE is_active END,
    image_url   = CASE WHEN image_url IS DISTINCT FROM $6 THEN $6 ELSE image_url END,
    price       = CASE WHEN price IS DISTINCT FROM $7 THEN $7 ELSE price END,
    stock_qty   = CASE WHEN stock_qty IS DISTINCT FROM $8 THEN $8 ELSE stock_qty END,
    attributes  = CASE WHEN attributes IS DISTINCT FROM $9 THEN $9 ELSE attributes END,
    updated_at  = CASE
                      WHEN
                          name IS DISTINCT FROM $2 OR
                          description IS DISTINCT FROM $3 OR
                          category_id IS DISTINCT FROM $4 OR
                          is_active IS DISTINCT FROM $5 OR
                          image_url IS DISTINCT FROM $6 OR
                          price IS DISTINCT FROM $7 OR
                          stock_qty IS DISTINCT FROM $8 OR
                          attributes IS DISTINCT FROM $9
                          THEN $10
                      ELSE updated_at END
WHERE id = $1;

-- name: DeleteProduct :exec
DELETE
FROM products
WHERE id = @product_id;


-- name: ListCategories :many
SELECT id, name
FROM categories
ORDER BY name;

-- name: CreateCategory :exec
INSERT INTO categories (id, name)
VALUES ($1, $2);

-- name: DeleteCategory :exec
DELETE
FROM categories
WHERE id = $1;

-- name: GetCategoryByID :one
SELECT *
FROM categories
WHERE id = $1;