package memory_repositories

import (
	"context"
	"kryptify/entities"
	appError "kryptify/usecases/errors"
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

func (repo *MemoryCredentialsRepository) DeleteCredential(ctx context.Context, credentialID string) error {
	repo.Lock()
	defer repo.Unlock()

	_, exists := repo.Credentials[credentialID]
	if !exists {
		return appError.NewErrorCredentialNotFound(credentialID)
	}

	delete(repo.Credentials, credentialID)
	return nil
}

func (repo *MemoryCredentialsRepository) UpdateCredential(ctx context.Context, credentialID string, updatedCredential *entities.Credential) error {
	repo.Lock()
	defer repo.Unlock()

	existingCredential, exists := repo.Credentials[credentialID]
	if !exists {
		return appError.NewErrorCredentialNotFound(credentialID)
	}

	if updatedCredential.Email != "" {
		existingCredential.Email = updatedCredential.Email
	}
	if updatedCredential.Category != "" {
		existingCredential.Category = updatedCredential.Category
	}
	if updatedCredential.Website != "" {
		existingCredential.Website = updatedCredential.Website
	}
	if updatedCredential.PasswordHash != "" {
		existingCredential.PasswordHash = updatedCredential.PasswordHash
	}

	repo.Credentials[credentialID] = existingCredential
	return nil
}

func (repo *MemoryCredentialsRepository) GetCredentialByID(ctx context.Context, credentialID string) (*entities.Credential, error) {
	repo.Lock()
	defer repo.Unlock()

	credential, exists := repo.Credentials[credentialID]

	if !exists {
		return nil, appError.NewErrorCredentialNotFound(credentialID)
	}

	return credential, nil
}

func (repo *MemoryCredentialsRepository) GetCredentialsByOwnerID(ctx context.Context, ownerID string) ([]*entities.Credential, error) {
	repo.Lock()
	defer repo.Unlock()

	var credentials []*entities.Credential

	for _, credential := range repo.Credentials {
		if credential.OwnerID == ownerID {
			credentials = append(credentials, credential)
		}
	}

	return credentials, nil
}
