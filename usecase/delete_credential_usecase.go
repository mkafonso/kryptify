package usecase

import (
	"context"
	"errors"
	"kryptify/repository"
	appError "kryptify/usecase/errors"
)

type DeleteCredentialRequest struct {
	TargetCredentialID, RequestedByAccountID string
}

type DeleteCredentialResponse struct {
}

type DeleteCredential struct {
	accountRepo    repository.AccountsRepositoryInterface
	credentialRepo repository.CredentialsRepositoryInterface
}

func NewDeleteCredential(
	accountRepo repository.AccountsRepositoryInterface,
	credentialRepo repository.CredentialsRepositoryInterface,
) *DeleteCredential {
	return &DeleteCredential{
		accountRepo:    accountRepo,
		credentialRepo: credentialRepo,
	}
}

func (d *DeleteCredential) Execute(ctx context.Context, data *DeleteCredentialRequest) (*DeleteCredentialResponse, error) {
	// check if requestedByAccountID exists
	account, err := d.accountRepo.GetAccountByID(ctx, data.RequestedByAccountID)
	if err != nil {
		return nil, appError.NewErrorAccountNotFound(data.RequestedByAccountID)
	}

	// check if credential exists
	credential, err := d.credentialRepo.GetCredentialByID(ctx, data.TargetCredentialID)
	if err != nil {
		return nil, appError.NewErrorCredentialNotFound(data.TargetCredentialID)
	}

	// check permission
	if credential.OwnerID != account.ID.String() {
		return nil, appError.NewErrorMissingPermission()
	}

	// delete the credential
	err = d.credentialRepo.DeleteCredential(ctx, data.TargetCredentialID)
	if err != nil {
		return nil, errors.New("error while deleting credential")
	}

	return &DeleteCredentialResponse{}, nil
}
