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
select feeds.name AS feed_name, feeds.url AS feed_url, users.name AS feed_creator_name
from feeds
INNER JOIN users ON feeds.user_id = users.id;
