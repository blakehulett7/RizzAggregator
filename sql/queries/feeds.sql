-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: NukeFeedsDB :exec
DELETE FROM feeds;

-- name: GetFeeds :many
SELECT * FROM feeds;
