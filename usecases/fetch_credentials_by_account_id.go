package usecases

import (
	"context"
	"errors"
	"kryptify/entities"
	"kryptify/repositories"
	appError "kryptify/usecases/errors"
)

type FetchCredentialsByAccountIDRequest struct {
	RequestedByAccountID string
}

type FetchCredentialsByAccountIDResponse struct {
	Credentials []*entities.Credential
}

type FetchCredentialsByAccountID struct {
	accountRepo    repositories.AccountsRepositoryInterface
	credentialRepo repositories.CredentialsRepositoryInterface
}

func NewFetchCredentialsByAccountID(
	accountRepo repositories.AccountsRepositoryInterface,
	credentialRepo repositories.CredentialsRepositoryInterface,
) *FetchCredentialsByAccountID {
	return &FetchCredentialsByAccountID{
		accountRepo:    accountRepo,
		credentialRepo: credentialRepo,
	}
}

func (f *FetchCredentialsByAccountID) Execute(ctx context.Context, data *FetchCredentialsByAccountIDRequest) (*FetchCredentialsByAccountIDResponse, error) {
	// check if requestedByAccountID exists
	_, err := f.accountRepo.GetAccountByID(ctx, data.RequestedByAccountID)
	if err != nil {
		return nil, appError.NewErrorAccountNotFound(data.RequestedByAccountID)
	}

	// Fetch credentials by account ID
	credentials, err := f.credentialRepo.GetCredentialsByOwnerID(ctx, data.RequestedByAccountID)
	if err != nil {
		return nil, errors.New("error while fetching credentials")
	}

	return &FetchCredentialsByAccountIDResponse{Credentials: credentials}, nil
}
