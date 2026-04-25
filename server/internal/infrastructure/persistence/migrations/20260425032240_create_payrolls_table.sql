-- +goose Up
CREATE TABLE payrolls (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    employee_id UUID NOT NULL REFERENCES employees(id),
    period_start DATE NOT NULL,
    period_end DATE NOT NULL,
    base_salary DECIMAL(12,2) NOT NULL,
    total_deduction DECIMAL(12,2) DEFAULT 0,  -- deduction/denda
    net_salary DECIMAL(12,2) NOT NULL,         -- base - deduction
    notes TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS payrolls;
