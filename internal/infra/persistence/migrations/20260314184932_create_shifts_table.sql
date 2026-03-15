-- +goose Up
CREATE TABLE shifts (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  opened_by UUID NOT NULL REFERENCES users(id),
  closed_by UUID NOT NULL REFERENCES users(id),
  opening_cash DECIMAL(12, 2) NOT NULL DEFAULT 0,
  closing_cash DECIMAL(12, 2),
  notes TEXT,
  opened_at TIMESTAMP NOT NULL DEFAULT NOW(),
  closed_at TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS shifts;