-- name: CreateEntry :one
INSERT INTO entries (
    account_id, amount
) VALUES (
 $1, $2
) RETURNING *;

-- name: GetEntry :one
SELECT id, account_id, amount, created_at FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT id, account_id, amount created_at FROM entries
WHERE account_id = $1
ORDER BY id
LIMIT $2 OFFSET $3;


-- name: DeleteEntry :exec
DELETE FROM entries WHERE id = $1 ;



