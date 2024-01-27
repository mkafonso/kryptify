package repository

import (
	"context"
	"kryptify/entity"
)

type AccountsRepositoryInterface interface {
	CreateAccount(ctx context.Context, account *entity.Account) error
	UpdateAccount(ctx context.Context, email string, updatedAccount *entity.Account) error
	GetAccountByID(ctx context.Context, accountID string) (*entity.Account, error)
	FindAccountByEmail(ctx context.Context, email string) (*entity.Account, error)
}
