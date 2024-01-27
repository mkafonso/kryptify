package memory_repository

import (
	"context"
	"kryptify/entity"
	"sync"
)

type MemorySessionsRepository struct {
	sync.Mutex
	Sessions map[string]*entity.Session // Map to store sessions, using sessionID as the key
}

func NewMemorySessionsRepository() *MemorySessionsRepository {
	return &MemorySessionsRepository{
		Sessions: make(map[string]*entity.Session),
	}
}

func (repo *MemorySessionsRepository) CreateSession(ctx context.Context, session *entity.Session) (*entity.Session, error) {
	repo.Lock()
	defer repo.Unlock()

	repo.Sessions[session.ID.String()] = session
	return session, nil
}
