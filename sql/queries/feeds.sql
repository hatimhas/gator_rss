-- name: CreateFeed :one
INSERT INTO feeds 
( id,
  created_at, 
  updated_at,
  url,
  name,
  user_id)
  VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6

  )
RETURNING *;

-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE url = $1;

-- name: GetAllFeeds :many
SELECT feeds.name AS feed_name, feeds.url AS feed_url, users.name AS feed_creator_name
FROM feeds
INNER JOIN users ON feeds.user_id = users.id;

-- name: MarkFeedAsFetched :exec
UPDATE feeds
SET 
  last_fetched_at = NOW(),
  updated_at = NOW()
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT url FROM feeds
ORDER BY LAST_FETCHED_AT ASC NULLS FIRST
LIMIT 1;
