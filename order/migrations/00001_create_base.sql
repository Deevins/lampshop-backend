-- +goose Up
CREATE TYPE payment_status AS ENUM ('pending', 'succeeded', 'failed');
CREATE TYPE payment_provider AS ENUM ('stripe', 'sbp');


CREATE TABLE orders (
                        id UUID PRIMARY KEY,
                        status payment_status NOT NULL,
                        total NUMERIC(10,2) NOT NULL CHECK (total >= 0),
                        is_active BOOLEAN DEFAULT true NOT NULL,
                        created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
                        updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

CREATE TABLE order_items (
                             id UUID PRIMARY KEY,
                             order_id UUID NOT NULL,
                             product_id UUID NOT NULL,
                             qty INTEGER NOT NULL CHECK (qty > 0),
                             unit_price NUMERIC(10,2) NOT NULL CHECK (unit_price >= 0)
);

CREATE TABLE payments (
                          id UUID PRIMARY KEY,
                          order_id UUID NOT NULL,
                          provider payment_provider NOT NULL,
                          status payment_status NOT NULL,
                          amount NUMERIC(10,2) NOT NULL CHECK (amount >= 0),
                          transaction_ref VARCHAR(120),
                          created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
                          updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS payments;
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;

DROP TYPE IF EXISTS payment_provder;
DROP TYPE IF EXISTS payment_status;
