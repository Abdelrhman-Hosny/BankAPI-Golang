-- name: CreateTransfer :one
INSERT INTO transfers (
    to_account_id,
    from_account_id,
    amount
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetTranfer :one
SELECT * FROM transfers
WHERE id = $1
LIMIT 1;

-- name: ListTranfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTransfer :exec
UPDATE transfers
SET amount = $2
WHERE id = $1;

-- name: DeleteTransfer :exec
DELETE FROM transfers WHERE id = $1;