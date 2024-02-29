// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: session.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createSession = `-- name: CreateSession :exec
INSERT INTO sessions (id, account_id, refresh_token, user_agent, client_ip, is_blocked, expires_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, account_id, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at
`

type CreateSessionParams struct {
	ID           uuid.UUID   `json:"id"`
	AccountID    string      `json:"account_id"`
	RefreshToken string      `json:"refresh_token"`
	UserAgent    pgtype.Text `json:"user_agent"`
	ClientIp     pgtype.Text `json:"client_ip"`
	IsBlocked    pgtype.Bool `json:"is_blocked"`
	ExpiresAt    time.Time   `json:"expires_at"`
}

// Create a new session
// Parameters: id, account_id, refresh_token, user_agent, client_ip, is_blocked, expires_at
// Returns: Newly created session
func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) error {
	_, err := q.db.Exec(ctx, createSession,
		arg.ID,
		arg.AccountID,
		arg.RefreshToken,
		arg.UserAgent,
		arg.ClientIp,
		arg.IsBlocked,
		arg.ExpiresAt,
	)
	return err
}
