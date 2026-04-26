-- +goose Up
DROP TABLE IF EXISTS expenses;
CREATE TABLE cash_transactions (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    coa_id      UUID NOT NULL,
    shift_id    UUID,
    type        VARCHAR(20) NOT NULL,
    amount      DECIMAL(12, 2) DEFAULT 0,
    description TEXT,
    date        DATE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL,

    CONSTRAINT fk_coa FOREIGN KEY (coa_id) REFERENCES chart_of_accounts(id),
    CONSTRAINT fk_shift FOREIGN KEY (shift_id) REFERENCES shifts(id)
);

CREATE INDEX idx_cash_transactions_deleted_at ON cash_transactions(deleted_at);
CREATE INDEX idx_cash_transactions_shift_id ON cash_transactions(shift_id);

-- +goose Down
DROP TABLE IF EXISTS cash_transactions;
