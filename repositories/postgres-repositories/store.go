package store

import (
	"database/sql"
	db "kryptify/db/sqlc"
)

type PostgresRepository struct {
	*db.Queries
	DB *sql.DB
}

func NewPostgresRepository(database *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		DB:      database,
		Queries: db.New(database),
	}
}
