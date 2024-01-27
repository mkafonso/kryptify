package usecase

import (
	"context"
	"kryptify/entity"
	"kryptify/repository"
	appError "kryptify/usecase/error"
)

type CreateCredentialRequest struct {
	RequestedByAccountID     string
	Email, Password, Website string
}

type CreateCredentialResponse struct {
	Credential *entity.Credential
}

type CreateCredential struct {
	accountRepo    repository.AccountsRepositoryInterface
	credentialRepo repository.CredentialsRepositoryInterface
}

func NewCreateCredential(
	accountRepo repository.AccountsRepositoryInterface,
	credentialRepo repository.CredentialsRepositoryInterface,
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
	credential, err := entity.NewCredential(data.Email, data.Password, data.Website, data.RequestedByAccountID)
	if err != nil {
		return nil, err
	}

	// save the new credential
	err = c.credentialRepo.CreateCredential(ctx, credential)
	if err != nil {
		return nil, err
	}

	response := &CreateCredentialResponse{
		Credential: &entity.Credential{
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
