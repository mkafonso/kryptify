-- Create a new account
-- Parameters: name, email, password_hash
-- Returns: Newly created account
-- name: CreateAccount :exec
INSERT INTO accounts (name, email, password_hash)
VALUES ($1, $2, $3)
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
-- Parameters: name, email, avatar_url, password_hash, updated_at, email
-- Returns: Updated account
-- name: UpdateAccount :exec
UPDATE accounts
SET
    name = $1,
    email = $2,
    avatar_url = $3,
    password_hash = $4,
    updated_at = $5
WHERE email = $6
RETURNING *;
