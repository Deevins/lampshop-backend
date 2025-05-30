-- +goose NO TRANSACTION
-- +goose Up

CREATE TABLE categories
(
    id   UUID PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE products
(
    id          UUID PRIMARY KEY,
    sku         VARCHAR(50) UNIQUE NOT NULL,
    name        VARCHAR(120)       NOT NULL,
    description TEXT,
    category_id UUID               NOT NULL,
    is_active   BOOLEAN                     DEFAULT true NOT NULL,
    image_url   TEXT,
    price       NUMERIC(10, 2)     NOT NULL CHECK (price > 0),
    stock_qty   INTEGER            NOT NULL CHECK (stock_qty >= 0),
    attributes  JSONB              NOT NULL,
    created_at  TIMESTAMPTZ        NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ        NOT NULL DEFAULT now()
);



-- +goose Down
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS products;