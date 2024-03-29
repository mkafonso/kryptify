package usecase

import (
	"context"

	"kryptify/entity"
	"kryptify/repository"
	appError "kryptify/usecase/error"
)

type CreateAccountRequest struct {
	Name, Email, Password string
}

type CreateAccountResponse struct {
	Account *entity.Account
}

type CreateAccount struct {
	accountRepo repository.AccountsRepositoryInterface
}

func NewCreateAccount(repo repository.AccountsRepositoryInterface) *CreateAccount {
	return &CreateAccount{accountRepo: repo}
}

func (c *CreateAccount) Execute(ctx context.Context, data *CreateAccountRequest) (*CreateAccountResponse, error) {
	// check if email is already taken
	foundAccount, _ := c.accountRepo.FindAccountByEmail(ctx, data.Email)
	if foundAccount != nil {
		return nil, appError.NewErrorEmailAlreadyTaken()
	}

	account, err := entity.NewAccount(data.Name, data.Email, data.Password)
	if err != nil {
		return nil, err
	}

	err = c.accountRepo.CreateAccount(ctx, account)
	if err != nil {
		return nil, err
	}

	return &CreateAccountResponse{Account: account}, nil
}
