package repository

import (
	"context"
	"kryptify/entity"
)

type SessionsRepositoryInterface interface {
	CreateSession(ctx context.Context, session *entity.Session) (*entity.Session, error)
}
