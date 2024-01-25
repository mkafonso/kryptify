package usecases

import (
	"context"
	"kryptify/entities"
	"kryptify/repositories"
)

type CreateCredentialRequest struct {
	Email, Password, Website, OwnerID string
}

type CreateCredentialResponse struct {
	Credential *entities.Credential
}

type CreateCredential struct {
	credentialRepo repositories.CredentialsRepositoryInterface
}

func NewCreateCredential(repo repositories.CredentialsRepositoryInterface) *CreateCredential {
	return &CreateCredential{credentialRepo: repo}
}

func (c *CreateCredential) Execute(ctx context.Context, data *CreateCredentialRequest) (*CreateCredentialResponse, error) {
	// create the new credential
	credential, err := entities.NewCredential(data.Email, data.Password, data.Website, data.OwnerID)
	if err != nil {
		return nil, err
	}

	// save the new credential
	err = c.credentialRepo.CreateCredential(ctx, credential)
	if err != nil {
		return nil, err
	}

	response := &CreateCredentialResponse{
		Credential: &entities.Credential{
			ID:        credential.ID,
			Email:     credential.Email,
			Website:   credential.Website,
			Category:  credential.Category,
			OwnerID:   credential.OwnerID,
			Health:    credential.Health,
			CreatedAt: credential.CreatedAt,
		},
	}

	return response, nil
}
