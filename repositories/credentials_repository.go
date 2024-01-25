package repositories

import (
	"context"
	"kryptify/entities"
)

type CredentialsRepositoryInterface interface {
	CreateCredential(ctx context.Context, credential *entities.Credential) error
	DeleteCredential(ctx context.Context, credentialID string) error
	UpdateCredential(ctx context.Context, credentialID string, updatedCredential *entities.Credential) error
	GetCredentialByID(ctx context.Context, credentialID string) (*entities.Credential, error)
	GetCredentialsByOwnerID(ctx context.Context, ownerID string) ([]*entities.Credential, error)
}
