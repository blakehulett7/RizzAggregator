-- name: CreateFeedFollows :one
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: NukeFeedFollowsDB :exec
DELETE FROM feed_follows;

-- name: GetFollows :many
SELECT * FROM feed_follows;

-- name: DeleteFollow :exec
DELETE FROM feed_follows
WHERE id = $1
AND user_id = $2;
