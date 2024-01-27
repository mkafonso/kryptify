package repository

import (
	"context"
	"kryptify/entities"
)

type SessionsRepositoryInterface interface {
	CreateSession(ctx context.Context, session *entities.Session) (*entities.Session, error)
}
