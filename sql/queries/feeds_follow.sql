-- name: CreateFeedsFollow :one
INSERT INTO feeds_follow (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFeedsFollow :many
SELECT * FROM feeds_follow where user_id=$1;

-- name: DeleteFeedsFollow :exec
DELETE FROM feeds_follow WHERE id=$1;
