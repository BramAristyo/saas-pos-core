-- +goose Up
CREATE TABLE payments (
  id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  order_id   UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
  method     VARCHAR(10) NOT NULL CHECK (method IN ('cash', 'qris', 'transfer', 'other')),
  amount     DECIMAL(12,2) NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS payments;
