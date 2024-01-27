package usecase

import (
	"context"
	"errors"
	"kryptify/entity"
	"kryptify/repository"
	appError "kryptify/usecase/error"
)

type FetchCredentialsByOwnerIDRequest struct {
	RequestedByAccountID string
}

type FetchCredentialsByOwnerIDResponse struct {
	Credentials []*entity.Credential
}

type FetchCredentialsByOwnerID struct {
	accountRepo    repository.AccountsRepositoryInterface
	credentialRepo repository.CredentialsRepositoryInterface
}

func NewFetchCredentialsByOwnerID(
	accountRepo repository.AccountsRepositoryInterface,
	credentialRepo repository.CredentialsRepositoryInterface,
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
