-- +goose Up
-- +goose StatementBegin
CREATE TABLE feeds
(
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  url TEXT UNIQUE NOT NULL,
  name TEXT UNIQUE NOT NULL,
  user_id UUID NOT NULL,
  CONSTRAINT fk_user
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE CASCADE
  );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
