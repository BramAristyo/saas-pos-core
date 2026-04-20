-- +goose Up
CREATE TABLE ledgers (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  coa_id UUID NOT NULL REFERENCES chart_of_accounts(id),
  shift_id UUID REFERENCES shifts(id),
  amount DECIMAL(15, 2) NOT NULL,
  notes TEXT,
  reference_id UUID,
  reference_type VARCHAR(50),
  transaction_date DATE NOT NULL,
  created_by UUID REFERENCES users(id),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS ledgers;
