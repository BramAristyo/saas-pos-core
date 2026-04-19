-- +goose Up
CREATE TABLE chart_of_accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    type VARCHAR(10) NOT NULL CHECK (type IN ('in', 'out')),
    is_system BOOLEAN DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

ALTER TABLE shift_expenses
    DROP COLUMN type,
    ADD COLUMN coa_id UUID NOT NULL REFERENCES chart_of_accounts(id);

ALTER TABLE expenses
    DROP COLUMN category,
    ADD COLUMN coa_id UUID NOT NULL REFERENCES chart_of_accounts(id);

INSERT INTO chart_of_accounts (name, type, is_system) VALUES ('Sales', 'in', true);

-- +goose Down
ALTER TABLE expenses
    DROP COLUMN coa_id,
    ADD COLUMN category VARCHAR(100) NOT NULL DEFAULT 'other';

ALTER TABLE shift_expenses
    DROP COLUMN coa_id,
    ADD COLUMN type VARCHAR(10) NOT NULL DEFAULT 'out' CHECK (type IN ('in', 'out'));

DROP TABLE IF EXISTS chart_of_accounts;
