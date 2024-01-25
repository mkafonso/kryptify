package usecases

import (
	"context"
	"kryptify/entities"
	"kryptify/repositories"
	appError "kryptify/usecases/errors"
)

type CreateCredentialRequest struct {
	RequestedByAccountID     string
	Email, Password, Website string
}

type CreateCredentialResponse struct {
	Credential *entities.Credential
}

type CreateCredential struct {
	accountRepo    repositories.AccountsRepositoryInterface
	credentialRepo repositories.CredentialsRepositoryInterface
}

func NewCreateCredential(
	accountRepo repositories.AccountsRepositoryInterface,
	credentialRepo repositories.CredentialsRepositoryInterface,
) *CreateCredential {
	return &CreateCredential{
		accountRepo:    accountRepo,
		credentialRepo: credentialRepo,
	}
}

func (c *CreateCredential) Execute(ctx context.Context, data *CreateCredentialRequest) (*CreateCredentialResponse, error) {
	// check if requestedByAccountID exists
	_, err := c.accountRepo.GetAccountByID(ctx, data.RequestedByAccountID)
	if err != nil {
		return nil, appError.NewErrorAccountNotFound(data.RequestedByAccountID)
	}

	// create the new credential
	credential, err := entities.NewCredential(data.Email, data.Password, data.Website, data.RequestedByAccountID)
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
