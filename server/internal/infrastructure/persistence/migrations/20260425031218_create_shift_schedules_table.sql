-- +goose Up
CREATE TABLE shift_schedules (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    tolerance_minutes INT NOT NULL DEFAULT 15,
    late_interval_minutes INT NOT NULL DEFAULT 10,
    late_deduction_amount DECIMAL(12,2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS shift_schedules;
