-- +goose Up
ALTER TABLE discounts ADD COLUMN is_custom_amount BOOLEAN NOT NULL DEFAULT false;

-- +goose Down
ALTER TABLE discounts DROP COLUMN is_custom_amount;
