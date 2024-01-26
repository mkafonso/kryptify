package repositories

import (
	"context"
	"kryptify/entities"
)

type AccountsRepositoryInterface interface {
	CreateAccount(ctx context.Context, account *entities.Account) error
	GetAccountByID(ctx context.Context, accountID string) (*entities.Account, error)
	FindAccountByEmail(ctx context.Context, email string) (*entities.Account, error)
}
