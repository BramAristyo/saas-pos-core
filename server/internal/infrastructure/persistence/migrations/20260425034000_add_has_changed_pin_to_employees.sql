-- +goose Up
ALTER TABLE employees ADD COLUMN IF NOT EXISTS has_changed_pin BOOLEAN DEFAULT FALSE;

-- +goose Down
ALTER TABLE employees DROP COLUMN IF EXISTS has_changed_pin;
