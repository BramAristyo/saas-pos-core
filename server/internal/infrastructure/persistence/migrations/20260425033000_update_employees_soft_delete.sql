-- +goose Up
ALTER TABLE employees DROP COLUMN IF EXISTS is_active;
ALTER TABLE employees ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP DEFAULT NULL;
CREATE INDEX IF NOT EXISTS idx_employees_deleted_at ON employees(deleted_at);

-- +goose Down
DROP INDEX IF EXISTS idx_employees_deleted_at;
ALTER TABLE employees DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE employees ADD COLUMN IF NOT EXISTS is_active BOOLEAN DEFAULT TRUE;
