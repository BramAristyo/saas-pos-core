-- +goose Up
CREATE TABLE bundling_packages (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(100) NOT NULL, 
  description TEXT,
  price DECIMAL(12, 2) NOT NULL DEFAULT 0,
  cogs DECIMAL(12, 2) NOT NULL DEFAULT 0,
  image_url VARCHAR(255),
  is_active BOOLEAN NOT NULL DEFAULT true,
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS bundling_packages;