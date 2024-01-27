package usecase

import (
	"context"
	"errors"
	"kryptify/entities"
	valueobjects "kryptify/entities/value-objects"
	"kryptify/repository"
	appError "kryptify/usecase/errors"
	"time"
)

type UpdateCredentialRequest struct {
	TargetCredentialID, RequestedByAccountID string
	Email, Website, Category, Password       string
}

type UpdateCredentialResponse struct {
	Credential *entities.Credential
}

type UpdateCredential struct {
	accountRepo    repository.AccountsRepositoryInterface
	credentialRepo repository.CredentialsRepositoryInterface
}

func NewUpdateCredential(
	accountRepo repository.AccountsRepositoryInterface,
	credentialRepo repository.CredentialsRepositoryInterface,
) *UpdateCredential {
	return &UpdateCredential{
		accountRepo:    accountRepo,
		credentialRepo: credentialRepo,
	}
}

func (u *UpdateCredential) Execute(ctx context.Context, data *UpdateCredentialRequest) (*UpdateCredentialResponse, error) {
	// check if requestedByAccountID exists
	account, err := u.accountRepo.GetAccountByID(ctx, data.RequestedByAccountID)
	if err != nil {
		return nil, appError.NewErrorAccountNotFound(data.RequestedByAccountID)
	}

	// check if credential exists
	credential, err := u.credentialRepo.GetCredentialByID(ctx, data.TargetCredentialID)
	if err != nil {
		return nil, appError.NewErrorCredentialNotFound(data.TargetCredentialID)
	}

	// check permission
	if credential.OwnerID != account.ID.String() {
		return nil, appError.NewErrorMissingPermission()
	}

	credentialUpdated := u.updateCredentialDetails(credential, data)
	err = u.credentialRepo.UpdateCredential(ctx, data.TargetCredentialID, credentialUpdated)
	if err != nil {
		return nil, errors.New("error while updating credential")
	}

	return &UpdateCredentialResponse{Credential: credentialUpdated}, nil
}

func (u *UpdateCredential) updateCredentialDetails(credential *entities.Credential, data *UpdateCredentialRequest) *entities.Credential {
	if data.Email != "" {
		credential.Email = data.Email
	}

	if data.Website != "" {
		credential.Website = data.Website
	}

	if data.Category != "" {
		credential.Category = data.Category
	}

	if data.Password != "" {
		credential.PasswordHash = valueobjects.NewPassword(data.Password)
	}

	credential.UpdatedAt = time.Now().UTC()

	return credential
}
