-- +goose Up
CREATE TABLE employees (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code VARCHAR(100) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    phone VARCHAR(20),
    base_salary DECIMAL(12,2) NOT NULL DEFAULT 0,

    pin_hash VARCHAR(255) NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE INDEX idx_employees_deleted_at ON employees(deleted_at);

-- +goose Down
DROP INDEX IF EXISTS idx_employees_deleted_at;
DROP TABLE IF EXISTS employees;
