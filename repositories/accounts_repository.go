package repositories

import (
	"context"
	"kryptify/entities"
)

type AccountsRepositoryInterface interface {
	GetAccountByID(ctx context.Context, accountID string) (*entities.Account, error)
}
