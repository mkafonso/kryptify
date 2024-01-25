package repositories

import (
	"context"
	"kryptify/entities"
)

type CredentialsRepositoryInterface interface {
	CreateCredential(ctx context.Context, credential *entities.Credential) error
	GetCredentialByID(ctx context.Context, credentialID string) (*entities.Credential, error)
}
