-- name: CreateAccount :one
INSERT INTO accouns (
    owner, balance, currency
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM accouns WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accouns ORDER BY id 
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE accouns 
SET balance = $1
WHERE id = $2
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accouns 
WHERE id = $1
RETURNING *;
