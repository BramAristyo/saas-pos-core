-- +goose Up
-- Refactor users
ALTER TABLE users ADD COLUMN deleted_at TIMESTAMP DEFAULT NULL;
CREATE INDEX idx_users_deleted_at ON users(deleted_at);
ALTER TABLE users DROP COLUMN is_active;

-- Refactor product_categories
ALTER TABLE product_categories ADD COLUMN deleted_at TIMESTAMP DEFAULT NULL;
CREATE INDEX idx_product_categories_deleted_at ON product_categories(deleted_at);
ALTER TABLE product_categories DROP COLUMN is_active;

-- Refactor products
ALTER TABLE products ADD COLUMN deleted_at TIMESTAMP DEFAULT NULL;
CREATE INDEX idx_products_deleted_at ON products(deleted_at);
ALTER TABLE products DROP COLUMN is_active;

-- Refactor modifier_groups
ALTER TABLE modifier_groups ADD COLUMN deleted_at TIMESTAMP DEFAULT NULL;
CREATE INDEX idx_modifier_groups_deleted_at ON modifier_groups(deleted_at);
ALTER TABLE modifier_groups DROP COLUMN is_active;

-- Refactor modifier_options
ALTER TABLE modifier_options ADD COLUMN deleted_at TIMESTAMP DEFAULT NULL;
CREATE INDEX idx_modifier_options_deleted_at ON modifier_options(deleted_at);
ALTER TABLE modifier_options DROP COLUMN is_active;

-- Refactor bundling_packages
ALTER TABLE bundling_packages ADD COLUMN deleted_at TIMESTAMP DEFAULT NULL;
CREATE INDEX idx_bundling_packages_deleted_at ON bundling_packages(deleted_at);
ALTER TABLE bundling_packages DROP COLUMN is_active;

-- Refactor discounts
ALTER TABLE discounts ADD COLUMN deleted_at TIMESTAMP DEFAULT NULL;
CREATE INDEX idx_discounts_deleted_at ON discounts(deleted_at);
ALTER TABLE discounts DROP COLUMN is_active;

-- Refactor taxes
ALTER TABLE taxes ADD COLUMN deleted_at TIMESTAMP DEFAULT NULL;
CREATE INDEX idx_taxes_deleted_at ON taxes(deleted_at);
ALTER TABLE taxes DROP COLUMN is_active;

-- Refactor sales_types
ALTER TABLE sales_types ADD COLUMN deleted_at TIMESTAMP DEFAULT NULL;
CREATE INDEX idx_sales_types_deleted_at ON sales_types(deleted_at);
ALTER TABLE sales_types DROP COLUMN is_active;

-- Refactor additional_charges
ALTER TABLE additional_charges ADD COLUMN deleted_at TIMESTAMP DEFAULT NULL;
CREATE INDEX idx_additional_charges_deleted_at ON additional_charges(deleted_at);
ALTER TABLE additional_charges DROP COLUMN is_active;

-- +goose Down
-- Restore additional_charges
DROP INDEX IF EXISTS idx_additional_charges_deleted_at;
ALTER TABLE additional_charges DROP COLUMN deleted_at;
ALTER TABLE additional_charges ADD COLUMN is_active BOOLEAN NOT NULL DEFAULT TRUE;

-- Restore sales_types
DROP INDEX IF EXISTS idx_sales_types_deleted_at;
ALTER TABLE sales_types DROP COLUMN deleted_at;
ALTER TABLE sales_types ADD COLUMN is_active BOOLEAN NOT NULL DEFAULT TRUE;

-- Restore taxes
DROP INDEX IF EXISTS idx_taxes_deleted_at;
ALTER TABLE taxes DROP COLUMN deleted_at;
ALTER TABLE taxes ADD COLUMN is_active BOOLEAN NOT NULL DEFAULT TRUE;

-- Restore discounts
DROP INDEX IF EXISTS idx_discounts_deleted_at;
ALTER TABLE discounts DROP COLUMN deleted_at;
ALTER TABLE discounts ADD COLUMN is_active BOOLEAN NOT NULL DEFAULT TRUE;

-- Restore bundling_packages
DROP INDEX IF EXISTS idx_bundling_packages_deleted_at;
ALTER TABLE bundling_packages DROP COLUMN deleted_at;
ALTER TABLE bundling_packages ADD COLUMN is_active BOOLEAN NOT NULL DEFAULT TRUE;

-- Restore modifier_options
DROP INDEX IF EXISTS idx_modifier_options_deleted_at;
ALTER TABLE modifier_options DROP COLUMN deleted_at;
ALTER TABLE modifier_options ADD COLUMN is_active BOOLEAN NOT NULL DEFAULT TRUE;

-- Restore modifier_groups
DROP INDEX IF EXISTS idx_modifier_groups_deleted_at;
ALTER TABLE modifier_groups DROP COLUMN deleted_at;
ALTER TABLE modifier_groups ADD COLUMN is_active BOOLEAN NOT NULL DEFAULT TRUE;

-- Restore products
DROP INDEX IF EXISTS idx_products_deleted_at;
ALTER TABLE products DROP COLUMN deleted_at;
ALTER TABLE products ADD COLUMN is_active BOOLEAN NOT NULL DEFAULT TRUE;

-- Restore product_categories
DROP INDEX IF EXISTS idx_product_categories_deleted_at;
ALTER TABLE product_categories DROP COLUMN deleted_at;
ALTER TABLE product_categories ADD COLUMN is_active BOOLEAN NOT NULL DEFAULT TRUE;

-- Restore users
DROP INDEX IF EXISTS idx_users_deleted_at;
ALTER TABLE users DROP COLUMN deleted_at;
ALTER TABLE users ADD COLUMN is_active BOOLEAN NOT NULL DEFAULT TRUE;
