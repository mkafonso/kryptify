package usecase

import (
	"context"
	"kryptify/entities"
	"kryptify/repository"
	appError "kryptify/usecase/errors"
)

type GetCredentialByIDRequest struct {
	TargetCredentialID, RequestedByAccountID string
}

type GetCredentialByIDResponse struct {
	Credential *entities.Credential
}

type GetCredentialByID struct {
	accountRepo    repository.AccountsRepositoryInterface
	credentialRepo repository.CredentialsRepositoryInterface
}

func NewGetCredentialByID(
	accountRepo repository.AccountsRepositoryInterface,
	credentialRepo repository.CredentialsRepositoryInterface,
) *GetCredentialByID {
	return &GetCredentialByID{
		accountRepo:    accountRepo,
		credentialRepo: credentialRepo,
	}
}

func (g *GetCredentialByID) Execute(ctx context.Context, data *GetCredentialByIDRequest) (*GetCredentialByIDResponse, error) {
	// check if requestedByAccountID exists
	account, err := g.accountRepo.GetAccountByID(ctx, data.RequestedByAccountID)
	if err != nil {
		return nil, appError.NewErrorAccountNotFound(data.RequestedByAccountID)
	}

	// check if credential exists
	credential, err := g.credentialRepo.GetCredentialByID(ctx, data.TargetCredentialID)
	if err != nil {
		return nil, appError.NewErrorCredentialNotFound(data.TargetCredentialID)
	}

	// check permission
	if credential.OwnerID != account.ID.String() {
		return nil, appError.NewErrorMissingPermission()
	}

	return &GetCredentialByIDResponse{Credential: credential}, nil
}
