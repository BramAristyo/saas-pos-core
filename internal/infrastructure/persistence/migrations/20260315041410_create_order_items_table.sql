-- +goose Up
CREATE TABLE order_items(
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  order_id        UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
  product_id      UUID REFERENCES products(id),
  bundling_id     UUID REFERENCES bundling_packages(id),
  discount_id     UUID REFERENCES discounts(id),
  product_name    VARCHAR(100) NOT NULL,
  product_price   DECIMAL(12,2) NOT NULL DEFAULT 0,
  product_cogs    DECIMAL(12,2) NOT NULL DEFAULT 0,
  quantity        INTEGER NOT NULL DEFAULT 1,
  discount_amount DECIMAL(12,2) NOT NULL DEFAULT 0,
  subtotal        DECIMAL(12,2) NOT NULL DEFAULT 0,
  created_at      TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at      TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS order_items;
