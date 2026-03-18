-- +goose Up
CREATE TABLE additional_charges (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sales_type_id UUID NOT NULL REFERENCES sales_types(id),
    name          VARCHAR(100) NOT NULL,
    type          VARCHAR(10) NOT NULL CHECK (type IN ('percentage', 'fixed')),
    amount        DECIMAL(12,2) NOT NULL DEFAULT 0,
    is_active     BOOLEAN NOT NULL DEFAULT true,
    created_at    TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS additional_charges;
