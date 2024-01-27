package usecase

import (
	"context"
	"kryptify/entities"
	"kryptify/repository"
	appError "kryptify/usecase/errors"
)

type GetAccountProfileRequest struct {
	TargetAccountEmail, RequestedByAccountID string
}

type GetAccountProfileResponse struct {
	Account *entities.Account
}

type GetAccountProfile struct {
	accountRepo repository.AccountsRepositoryInterface
}

func NewGetAccountProfile(repo repository.AccountsRepositoryInterface) *GetAccountProfile {
	return &GetAccountProfile{accountRepo: repo}
}

func (c *GetAccountProfile) Execute(ctx context.Context, data *GetAccountProfileRequest) (*GetAccountProfileResponse, error) {
	// check if requestedByAccountID exists
	account, err := c.accountRepo.GetAccountByID(ctx, data.RequestedByAccountID)
	if err != nil {
		return nil, appError.NewErrorAccountNotFound(data.RequestedByAccountID)
	}

	// check permission
	if account.Email != data.TargetAccountEmail {
		return nil, appError.NewErrorMissingPermission()
	}

	return &GetAccountProfileResponse{Account: account}, nil
}
