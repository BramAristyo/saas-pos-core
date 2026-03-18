-- +goose Up
CREATE TABLE bundling_items (
  id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  bundling_package_id UUID NOT NULL REFERENCES bundling_packages(id) ON DELETE CASCADE,
  product_id UUID NOT NULL REFERENCES products(id),
  qty INTEGER NOT NULL DEFAULT 1,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  UNIQUE (bundling_package_id, product_id)
);

-- +goose Down
DROP TABLE IF EXISTS bundling_items;
