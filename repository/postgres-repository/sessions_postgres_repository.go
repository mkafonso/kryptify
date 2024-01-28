package store

import (
	"context"
	db "kryptify/db/sqlc"
	"kryptify/entity"

	"github.com/jackc/pgx/v5/pgtype"
)

func (r *PostgresRepository) CreateSession(ctx context.Context, session *entity.Session) (*entity.Session, error) {
	params := db.CreateSessionParams{
		AccountID:    session.AccountID,
		RefreshToken: session.RefreshToken,
		UserAgent:    pgtype.Text{String: session.UserAgent, Valid: session.UserAgent != ""},
		ClientIp:     pgtype.Text{String: session.ClientIP, Valid: session.ClientIP != ""},
		IsBlocked:    pgtype.Bool{Bool: session.IsBlocked, Valid: true},
		ExpiresAt:    session.ExpiresAt,
	}

	err := r.Queries.CreateSession(ctx, params)
	return nil, err
}
