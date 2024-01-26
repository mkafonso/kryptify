package store

import (
	"context"
	"database/sql"
	db "kryptify/database/sqlc"
	"kryptify/entities"
)

func (r *PostgresRepository) CreateSession(ctx context.Context, session *entities.Session) (*entities.Session, error) {
	params := db.CreateSessionParams{
		AccountID:    session.AccountID,
		RefreshToken: session.RefreshToken,
		UserAgent:    sql.NullString{String: session.UserAgent, Valid: session.UserAgent != ""},
		ClientIp:     sql.NullString{String: session.ClientIP, Valid: session.ClientIP != ""},
		IsBlocked:    sql.NullBool{Bool: session.IsBlocked, Valid: true},
		ExpiresAt:    session.ExpiresAt,
	}

	err := r.Queries.CreateSession(ctx, params)
	return nil, err
}
