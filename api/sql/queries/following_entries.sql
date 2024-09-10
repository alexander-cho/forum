-- name: CreateFollowingEntry :one
INSERT INTO following_entries (id, created_at, updated_at, user_id, entry_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFollowingEntries :many
SELECT * FROM following_entries WHERE user_id=$1;

-- name: DeleteFollowingEntry :exec
DELETE FROM following_entries WHERE id=$1 AND user_id=$2;
