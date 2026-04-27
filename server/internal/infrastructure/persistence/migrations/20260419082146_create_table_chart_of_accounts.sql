-- +goose Up
CREATE TABLE chart_of_accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    type VARCHAR(10) NOT NULL CHECK (type IN ('in', 'out')),
    is_system BOOLEAN DEFAULT false,
    is_operational BOOLEAN DEFAULT true,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

INSERT INTO chart_of_accounts (name, type, is_system) VALUES ('Sales', 'in', true);

-- +goose Down
DROP TABLE IF EXISTS chart_of_accounts CASCADE;
