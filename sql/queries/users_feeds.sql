-- name: CreateFeedFollow :one
WITH inserted_feed_follow_in_users_feeds AS (
  INSERT INTO users_feeds (id, created_at, updated_at, user_id, feed_id) VALUES ($1, $2, $3,$4, $5)
  RETURNING *
) 
SELECT iff.*, u.name AS user_name, f.name AS feed_name
FROM inserted_feed_follow_in_users_feeds iff
INNER JOIN feeds f ON iff.feed_id = f.id
INNER JOIN users u ON iff.user_id = u.id;



