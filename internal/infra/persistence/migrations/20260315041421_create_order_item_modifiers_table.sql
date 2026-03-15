-- +goose Up
CREATE TABLE order_item_modifiers (
  id                 UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  order_item_id      UUID NOT NULL REFERENCES order_items(id) ON DELETE CASCADE,
  modifier_option_id UUID REFERENCES modifier_options(id),
  modifier_name      VARCHAR(100) NOT NULL,
  price_adjustment   DECIMAL(12,2) NOT NULL DEFAULT 0,
  cogs_adjustment    DECIMAL(12,2) NOT NULL DEFAULT 0,
  created_at         TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS order_item_modifiers;
