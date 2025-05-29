-- name: CreateFeedFollow :one
WITH inserted_feed_follow_in_users_feeds AS (
  INSERT INTO users_feeds (id, created_at, updated_at, user_id, feed_id) VALUES ($1, $2, $3,$4, $5)
  RETURNING *
) 
SELECT iff.*, u.name AS user_name, f.name AS feed_name
FROM inserted_feed_follow_in_users_feeds iff
INNER JOIN feeds f ON iff.feed_id = f.id
INNER JOIN users u ON iff.user_id = u.id;


-- GetFeedFollowsForUser query. It should return all the feed follows for a given user, and include the names of the feeds and user in the result.
-- name: GetFeedFollowsForUser :many
SELECT uf.*, u.name AS user_name, f.name AS feed_name
FROM users_feeds uf
INNER JOIN users u ON uf.user_id = u.id
INNER JOIN feeds f ON uf.feed_id = f.id
WHERE u.name = $1;  



-- name: DeleteFeedFollow :exec
DELETE from users_feeds
WHERE users_feeds.user_id = $1  AND users_feeds.feed_id = ( SELECT id FROM feeds WHERE url =  $2); 
