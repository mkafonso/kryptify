package memory_repositories

import (
	"context"
	"kryptify/entities"
	"sync"
)

type MemoryCredentialsRepository struct {
	sync.Mutex
	Credentials map[string]*entities.Credential // Map to store credentials, using CredentialID as the key
}

func NewMemoryCredentialsRepository() *MemoryCredentialsRepository {
	return &MemoryCredentialsRepository{
		Credentials: make(map[string]*entities.Credential),
	}
}

func (repo *MemoryCredentialsRepository) CreateCredential(ctx context.Context, credential *entities.Credential) error {
	repo.Lock()
	defer repo.Unlock()

	repo.Credentials[credential.ID.String()] = credential
	return nil
}
