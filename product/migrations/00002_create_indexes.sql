-- +goose NO TRANSACTION
-- +goose Up

CREATE INDEX idx_products_name_ft ON products
    USING GIN (to_tsvector('simple', name));

CREATE INDEX idx_products_category ON products(category_id);


-- +goose Down

DROP INDEX IF EXISTS idx_products_name_ft;
DROP INDEX IF EXISTS idx_products_category;