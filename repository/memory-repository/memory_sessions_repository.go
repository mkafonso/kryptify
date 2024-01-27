package memory_repository

import (
	"context"
	"kryptify/entities"
	"sync"
)

type MemorySessionsRepository struct {
	sync.Mutex
	Sessions map[string]*entities.Session // Map to store sessions, using sessionID as the key
}

func NewMemorySessionsRepository() *MemorySessionsRepository {
	return &MemorySessionsRepository{
		Sessions: make(map[string]*entities.Session),
	}
}

func (repo *MemorySessionsRepository) CreateSession(ctx context.Context, session *entities.Session) (*entities.Session, error) {
	repo.Lock()
	defer repo.Unlock()

	repo.Sessions[session.ID.String()] = session
	return session, nil
}
