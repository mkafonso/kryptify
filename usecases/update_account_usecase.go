package usecases

import (
	"context"
	"errors"
	"kryptify/entities"
	valueobjects "kryptify/entities/value-objects"
	"kryptify/repositories"
	appError "kryptify/usecases/errors"
	"time"
)

type UpdateAccountRequest struct {
	TargetAccountID, RequestedByAccountID string
	Name, Email, AvatarUrl, Password      string
}

type UpdateAccountResponse struct {
	Account *entities.Account
}

type UpdateAccount struct {
	accountRepo repositories.AccountsRepositoryInterface
}

func NewUpdateAccount(repo repositories.AccountsRepositoryInterface) *UpdateAccount {
	return &UpdateAccount{accountRepo: repo}
}

func (c *UpdateAccount) Execute(ctx context.Context, data *UpdateAccountRequest) (*UpdateAccountResponse, error) {
	// check if requestedByAccountID exists
	account, err := c.accountRepo.GetAccountByID(ctx, data.RequestedByAccountID)
	if err != nil {
		return nil, appError.NewErrorAccountNotFound(data.RequestedByAccountID)
	}

	// check permission
	if data.TargetAccountID != data.RequestedByAccountID {
		return nil, appError.NewErrorMissingPermission()
	}

	accountUpdated := c.updateAccountDetails(account, data)
	err = c.accountRepo.UpdateAccount(ctx, account.Email, accountUpdated)
	if err != nil {
		return nil, errors.New("error while updating account")
	}

	return &UpdateAccountResponse{Account: accountUpdated}, nil
}

func (c *UpdateAccount) updateAccountDetails(account *entities.Account, data *UpdateAccountRequest) *entities.Account {
	if data.Name != "" {
		account.Name = data.Name
	}

	if data.Email != "" {
		account.Email = data.Email
	}

	if data.AvatarUrl != "" {
		account.AvatarUrl = data.AvatarUrl
	}

	if data.Password != "" {
		account.PasswordHash = valueobjects.NewPassword(data.Password)
	}

	account.UpdatedAt = time.Now().UTC()

	return account
}
