-- +goose Up
ALTER TABLE taxes ADD COLUMN status BOOLEAN DEFAULT true;

-- +goose Down
ALTER TABLE taxes DROP COLUMN status;
