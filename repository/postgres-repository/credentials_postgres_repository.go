package store

import (
	"context"
	db "kryptify/db/sqlc"
	"kryptify/entity"
	valueobject "kryptify/entity/value-object"
	"kryptify/util"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *PostgresRepository) CreateCredential(ctx context.Context, credential *entity.Credential) error {
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

func (r *PostgresRepository) UpdateCredential(ctx context.Context, credentialID string, updatedCredential *entity.Credential) error {
	params := db.UpdateCredentialParams{
		ID:           updatedCredential.ID,
		Email:        updatedCredential.Email,
		Website:      updatedCredential.Website,
		Category:     pgtype.Text{String: updatedCredential.Category, Valid: updatedCredential.Category != ""},
		PasswordHash: string(updatedCredential.PasswordHash),
		UpdatedAt:    updatedCredential.UpdatedAt,
	}

	err := r.Queries.UpdateCredential(ctx, params)
	return err
}

func (r *PostgresRepository) GetCredentialByID(ctx context.Context, credentialID string) (*entity.Credential, error) {
	credential, err := r.Queries.GetCredentialByID(ctx, uuid.MustParse(credentialID))
	if err != nil {
		return nil, err
	}

	entityCredential := &entity.Credential{
		ID:           credential.ID,
		Email:        credential.Email,
		Website:      credential.Website,
		Category:     util.GetStringValue(credential.Category),
		OwnerID:      credential.OwnerID.String(),
		PasswordHash: valueobject.Password(credential.PasswordHash),
		Health:       valueobject.Health(credential.Health),
		CreatedAt:    credential.CreatedAt,
		UpdatedAt:    credential.UpdatedAt,
	}

	return entityCredential, nil
}

func (r *PostgresRepository) GetCredentialsByOwnerID(ctx context.Context, ownerID string) ([]*entity.Credential, error) {
	credentials, err := r.Queries.GetCredentialsByOwnerID(ctx, uuid.MustParse(ownerID))
	if err != nil {
		return nil, err
	}

	var entityCredentials []*entity.Credential
	for _, credential := range credentials {
		entityCredential := &entity.Credential{
			ID:           credential.ID,
			Email:        credential.Email,
			Website:      credential.Website,
			Category:     util.GetStringValue(credential.Category),
			OwnerID:      credential.OwnerID.String(),
			PasswordHash: valueobject.Password(credential.PasswordHash),
			Health:       valueobject.Health(credential.Health),
			CreatedAt:    credential.CreatedAt,
			UpdatedAt:    credential.UpdatedAt,
		}
		entityCredentials = append(entityCredentials, entityCredential)
	}

	return entityCredentials, nil
}
