-- +goose Up
CREATE TABLE shift_expenses (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  shift_id UUID NOT NULL REFERENCES shifts(id),
  type VARCHAR(10) NOT NULL CHECK (type IN ('in', 'out')),
  amount DECIMAL(12, 2) NOT NULL DEFAULT 0,
  description TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS shift_expenses;
