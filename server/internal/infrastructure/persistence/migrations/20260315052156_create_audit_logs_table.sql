-- +goose Up
CREATE TABLE audit_logs (
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id     UUID NOT NULL,
  action      VARCHAR(50) NOT NULL,
  entity      VARCHAR(50) NOT NULL,
  entity_id   UUID,
  description TEXT,
  created_at  TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS audit_logs;
