package repository

import (
	"context"
	"kryptify/entity"
)

type CredentialsRepositoryInterface interface {
	CreateCredential(ctx context.Context, credential *entity.Credential) error
	DeleteCredential(ctx context.Context, credentialID string) error
	UpdateCredential(ctx context.Context, credentialID string, updatedCredential *entity.Credential) error
	GetCredentialByID(ctx context.Context, credentialID string) (*entity.Credential, error)
	GetCredentialsByOwnerID(ctx context.Context, ownerID string) ([]*entity.Credential, error)
}
