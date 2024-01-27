package store

import (
	"context"
	"database/sql"
	db "kryptify/db/sqlc"
	"kryptify/entity"
)

func (r *PostgresRepository) CreateSession(ctx context.Context, session *entity.Session) (*entity.Session, error) {
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
