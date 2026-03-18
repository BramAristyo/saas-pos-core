-- +goose Up
CREATE TABLE orders(
  ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  shift_id UUID NOT NULL REFERENCES shifts(id),
  cashier_id UUID NOT NULL REFERENCES users(id),
  sales_type_id UUID NOT NULL REFERENCES sales_types(id),
  tax_id UUID REFERENCES taxes(id),
  discount_id UUID REFERENCES discounts(id),

  order_number VARCHAR(50) NOT NULL UNIQUE,
  subtotal DECIMAL(12,2) NOT NULL DEFAULT 0,
  discount_amount DECIMAL(12, 2) NOT NULL DEFAULT 0,
  tax_amount DECIMAL(12, 2) NOT NULL DEFAULT 0,
  charge_amount DECIMAL(12,2) NOT NULL DEFAULT 0,
  total DECIMAL(12,2) NOT NULL DEFAULT 0,
  status VARCHAR(10) NOT NULL DEFAULT 'completed' CHECK (status IN ('completed', 'voided')),
  void_reason     TEXT,
  voided_by       UUID REFERENCES users(id),
  voided_at       TIMESTAMP,
  created_at      TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at      TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS orders;