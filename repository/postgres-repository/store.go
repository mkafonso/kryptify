package store

import (
	db "kryptify/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	*db.Queries
	connectionPool *pgxpool.Pool
}

func NewPostgresRepository(connectionPool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{
		connectionPool: connectionPool,
		Queries:        db.New(connectionPool),
	}
}
