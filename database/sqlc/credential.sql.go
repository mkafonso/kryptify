// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: credential.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createCredential = `-- name: CreateCredential :exec
INSERT INTO credentials (email, password_hash, website, owner_id)
VALUES ($1, $2, $3, $4)
RETURNING id, email, website, category, owner_id, password_hash, health, created_at, updated_at
`

type CreateCredentialParams struct {
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	Website      string    `json:"website"`
	OwnerID      uuid.UUID `json:"owner_id"`
}

// Create a new credential
// Parameters: email, password_hash, website, owner_id
// Returns: Newly created credential
func (q *Queries) CreateCredential(ctx context.Context, arg CreateCredentialParams) error {
	_, err := q.db.ExecContext(ctx, createCredential,
		arg.Email,
		arg.PasswordHash,
		arg.Website,
		arg.OwnerID,
	)
	return err
}

const deleteCredential = `-- name: DeleteCredential :exec
DELETE FROM credentials WHERE id = $1
`

// Delete a credential by ID
// Parameters: credentialID
func (q *Queries) DeleteCredential(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteCredential, id)
	return err
}

const getCredentialByID = `-- name: GetCredentialByID :one
SELECT id, email, website, category, owner_id, password_hash, health, created_at, updated_at FROM credentials
WHERE id = $1
LIMIT 1
`

// Get a credential by ID
// Parameters: id
// Returns: Credential with the specified ID
func (q *Queries) GetCredentialByID(ctx context.Context, id uuid.UUID) (Credential, error) {
	row := q.db.QueryRowContext(ctx, getCredentialByID, id)
	var i Credential
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Website,
		&i.Category,
		&i.OwnerID,
		&i.PasswordHash,
		&i.Health,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCredentialsByOwnerID = `-- name: GetCredentialsByOwnerID :many
SELECT id, email, website, category, owner_id, password_hash, health, created_at, updated_at FROM credentials
WHERE owner_id = $1
`

// Get all credentials belonging to a specific owner
// Parameters: ownerID
// Returns: List of credentials owned by the specified ID
func (q *Queries) GetCredentialsByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]Credential, error) {
	rows, err := q.db.QueryContext(ctx, getCredentialsByOwnerID, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Credential
	for rows.Next() {
		var i Credential
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Website,
			&i.Category,
			&i.OwnerID,
			&i.PasswordHash,
			&i.Health,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCredential = `-- name: UpdateCredential :exec
UPDATE credentials
SET
    email = $1,
    website = $2,
    category = $3,
    password_hash = $4
WHERE id = $5
RETURNING id, email, website, category, owner_id, password_hash, health, created_at, updated_at
`

type UpdateCredentialParams struct {
	Email        string         `json:"email"`
	Website      string         `json:"website"`
	Category     sql.NullString `json:"category"`
	PasswordHash string         `json:"password_hash"`
	ID           uuid.UUID      `json:"id"`
}

// Update a credential by ID
// Parameters: email, website, category, password_hash, id
// Returns: Updated credential
func (q *Queries) UpdateCredential(ctx context.Context, arg UpdateCredentialParams) error {
	_, err := q.db.ExecContext(ctx, updateCredential,
		arg.Email,
		arg.Website,
		arg.Category,
		arg.PasswordHash,
		arg.ID,
	)
	return err
}
