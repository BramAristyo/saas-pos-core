-- +goose Up
CREATE TABLE attendances (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    employee_id UUID NOT NULL REFERENCES employees(id),
    date DATE NOT NULL,
    check_in TIMESTAMP,
    check_out TIMESTAMP,
    late_minutes INT DEFAULT 0,             -- calculate
    deduction_amount DECIMAL(12,2) DEFAULT 0, -- calculate
    shift_schedule_id INT REFERENCES shift_schedules(id),
    notes TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(employee_id, date, shift_schedule_id)               -- record
);

-- +goose Down
DROP TABLE IF EXISTS attendances;
