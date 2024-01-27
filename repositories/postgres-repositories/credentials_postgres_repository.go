package store

import (
	"context"
	"database/sql"
	db "kryptify/db/sqlc"
	"kryptify/entities"
	valueobjects "kryptify/entities/value-objects"
	"kryptify/utils"

	"github.com/google/uuid"
)

func (r *PostgresRepository) CreateCredential(ctx context.Context, credential *entities.Credential) error {
	params := db.CreateCredentialParams{
		Email:        credential.Email,
		Website:      credential.Website,
		OwnerID:      uuid.MustParse(credential.OwnerID),
		PasswordHash: string(credential.PasswordHash),
	}

	err := r.Queries.CreateCredential(ctx, params)
	return err
}

func (r *PostgresRepository) DeleteCredential(ctx context.Context, credentialID string) error {
	err := r.Queries.DeleteCredential(ctx, uuid.MustParse(credentialID))
	return err
}

func (r *PostgresRepository) UpdateCredential(ctx context.Context, credentialID string, updatedCredential *entities.Credential) error {
	params := db.UpdateCredentialParams{
		ID:           updatedCredential.ID,
		Email:        updatedCredential.Email,
		Website:      updatedCredential.Website,
		Category:     sql.NullString{String: updatedCredential.Category, Valid: updatedCredential.Category != ""},
		PasswordHash: string(updatedCredential.PasswordHash),
		UpdatedAt:    updatedCredential.UpdatedAt,
	}

	err := r.Queries.UpdateCredential(ctx, params)
	return err
}

func (r *PostgresRepository) GetCredentialByID(ctx context.Context, credentialID string) (*entities.Credential, error) {
	credential, err := r.Queries.GetCredentialByID(ctx, uuid.MustParse(credentialID))
	if err != nil {
		return nil, err
	}

	entityCredential := &entities.Credential{
		ID:           credential.ID,
		Email:        credential.Email,
		Website:      credential.Website,
		Category:     utils.GetStringValue(credential.Category),
		OwnerID:      credential.OwnerID.String(),
		PasswordHash: valueobjects.Password(credential.PasswordHash),
		Health:       valueobjects.Health(credential.Health),
		CreatedAt:    credential.CreatedAt,
		UpdatedAt:    credential.UpdatedAt,
	}

	return entityCredential, nil
}

func (r *PostgresRepository) GetCredentialsByOwnerID(ctx context.Context, ownerID string) ([]*entities.Credential, error) {
	credentials, err := r.Queries.GetCredentialsByOwnerID(ctx, uuid.MustParse(ownerID))
	if err != nil {
		return nil, err
	}

	var entityCredentials []*entities.Credential
	for _, credential := range credentials {
		entityCredential := &entities.Credential{
			ID:           credential.ID,
			Email:        credential.Email,
			Website:      credential.Website,
			Category:     utils.GetStringValue(credential.Category),
			OwnerID:      credential.OwnerID.String(),
			PasswordHash: valueobjects.Password(credential.PasswordHash),
			Health:       valueobjects.Health(credential.Health),
			CreatedAt:    credential.CreatedAt,
			UpdatedAt:    credential.UpdatedAt,
		}
		entityCredentials = append(entityCredentials, entityCredential)
	}

	return entityCredentials, nil
}
