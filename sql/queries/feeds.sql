-- name: CreateFeed :one
INSERT INTO feeds (id, user_id, created_at, updated_at, last_fetched_at, name, url)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: NukeFeedsDB :exec
DELETE FROM feeds;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at NULLS FIRST
LIMIT $1;
