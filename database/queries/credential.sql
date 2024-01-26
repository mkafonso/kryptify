-- Create a new credential
-- Parameters: email, password_hash, website, owner_id
-- Returns: Newly created credential
-- name: CreateCredential :exec
INSERT INTO credentials (email, password_hash, website, owner_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- Delete a credential by ID
-- Parameters: credentialID
-- name: DeleteCredential :exec
DELETE FROM credentials WHERE id = $1;

-- Update a credential by ID
-- Parameters: email, website, category, password_hash, id
-- Returns: Updated credential
-- name: UpdateCredential :exec
UPDATE credentials
SET
    email = $1,
    website = $2,
    category = $3,
    password_hash = $4,
WHERE id = $5
RETURNING *;

-- Get a credential by ID
-- Parameters: id
-- Returns: Credential with the specified ID
-- name: GetCredentialByID :one
SELECT * FROM credentials
WHERE id = $1
LIMIT 1;

-- Get all credentials belonging to a specific owner
-- Parameters: ownerID
-- Returns: List of credentials owned by the specified ID
-- name: GetCredentialsByOwnerID :many
SELECT * FROM credentials
WHERE owner_id = $1;
