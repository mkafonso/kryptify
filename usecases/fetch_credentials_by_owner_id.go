package usecases

import (
	"context"
	"errors"
	"kryptify/entities"
	"kryptify/repositories"
	appError "kryptify/usecases/errors"
)

type FetchCredentialsByOwnerIDRequest struct {
	RequestedByAccountID string
}

type FetchCredentialsByOwnerIDResponse struct {
	Credentials []*entities.Credential
}

type FetchCredentialsByOwnerID struct {
	accountRepo    repositories.AccountsRepositoryInterface
	credentialRepo repositories.CredentialsRepositoryInterface
}

func NewFetchCredentialsByOwnerID(
	accountRepo repositories.AccountsRepositoryInterface,
	credentialRepo repositories.CredentialsRepositoryInterface,
) *FetchCredentialsByOwnerID {
	return &FetchCredentialsByOwnerID{
		accountRepo:    accountRepo,
		credentialRepo: credentialRepo,
	}
}

func (f *FetchCredentialsByOwnerID) Execute(ctx context.Context, data *FetchCredentialsByOwnerIDRequest) (*FetchCredentialsByOwnerIDResponse, error) {
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

	return &FetchCredentialsByOwnerIDResponse{Credentials: credentials}, nil
}
