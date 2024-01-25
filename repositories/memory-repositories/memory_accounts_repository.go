package memory_repositories

import (
	"context"
	"kryptify/entities"
	appError "kryptify/usecases/errors"
	"sync"
)

type MemoryAccountsRepository struct {
	sync.Mutex
	Accounts map[string]*entities.Account // Map to store accounts, using email as the key
}

func NewMemoryAccountsRepository() *MemoryAccountsRepository {
	return &MemoryAccountsRepository{
		Accounts: make(map[string]*entities.Account),
	}
}

func (repo *MemoryAccountsRepository) CreateAccount(ctx context.Context, account *entities.Account) error {
	repo.Lock()
	defer repo.Unlock()

	if _, exists := repo.Accounts[account.Email]; exists {
		err := appError.NewErrorEmailAlreadyTaken()
		return err
	}

	repo.Accounts[account.Email] = account
	return nil
}

func (repo *MemoryAccountsRepository) GetAccountByID(ctx context.Context, accountID string) (*entities.Account, error) {
	repo.Lock()
	defer repo.Unlock()

	for _, account := range repo.Accounts {
		if account.ID.String() == accountID {
			return account, nil
		}
	}

	return nil, appError.NewErrorAccountNotFound(accountID)
}
