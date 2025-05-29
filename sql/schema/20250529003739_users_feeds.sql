-- +goose Up
-- +goose StatementBegin
CREATE TABLE users_feeds
(
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  user_id UUID NOT NULL,
  feed_id UUID NOT NULL,

  CONSTRAINT fk_user_id_users
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE CASCADE,
  

  CONSTRAINT fk_feed_id_feeds
    FOREIGN KEY (feed_id)
    REFERENCES feeds (id)
    ON DELETE CASCADE,

  CONSTRAINT unique_user_feed
    UNIQUE (user_id, feed_id));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
