-- +goose Up
CREATE TABLE products (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  category_id UUID REFERENCES product_categories(id) ON DELETE SET NULL,
  name VARCHAR(100) NOT NULL UNIQUE,
  description TEXT,
  price DECIMAL(12, 2) NOT NULL DEFAULT 0,
  cogs DECIMAL(12, 2) NOT NULL DEFAULT 0,
  image_url VARCHAR(255),
  is_active BOOLEAN NOT NULL DEFAULT true,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS products;