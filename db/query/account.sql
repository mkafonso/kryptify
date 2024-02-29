-- Create a new account
-- Parameters: id, name, email, password_hash
-- Returns: Newly created account
-- name: CreateAccount :exec
INSERT INTO accounts (id, name, email, password_hash)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- Find an account by email
-- Parameters: email
-- Returns: Single account matching the email
-- name: FindAccountByEmail :one
SELECT * FROM accounts
WHERE email = $1
LIMIT 1;

-- Get an account by ID
-- Parameters: id
-- Returns: Single account matching the ID
-- name: GetAccountByID :one
SELECT * FROM accounts
WHERE id = $1
LIMIT 1;

-- Update an account by email
-- Parameters: name, avatar_url, password_hash, updated_at, email
-- Returns: Updated account
-- name: UpdateAccount :exec
UPDATE accounts
SET
    name = $1,
    avatar_url = $2,
    password_hash = $3,
    updated_at = $4
WHERE email = $5
RETURNING *;
