package usecase

import (
	"context"
	"errors"
	"kryptify/entity"
	"kryptify/repository"
	appError "kryptify/usecase/error"
	"kryptify/util"
	"time"
)

type UpdateAccountRequest struct {
	TargetAccountID, RequestedByAccountID string
	Name, AvatarUrl, Password             string
}

type UpdateAccountResponse struct {
	Account *entity.Account
}

type UpdateAccount struct {
	accountRepo repository.AccountsRepositoryInterface
}

func NewUpdateAccount(repo repository.AccountsRepositoryInterface) *UpdateAccount {
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

func (c *UpdateAccount) updateAccountDetails(account *entity.Account, data *UpdateAccountRequest) *entity.Account {
	if data.Name != "" {
		account.Name = data.Name
	}

	if data.AvatarUrl != "" {
		account.AvatarUrl = data.AvatarUrl
	}

	if data.Password != "" {
		account.PasswordHash = util.NewPassword(data.Password)
	}

	account.UpdatedAt = time.Now().UTC()

	return account
}
