package memory_repository

import (
	"context"
	"kryptify/entity"
	appError "kryptify/usecase/error"
	"sync"
)

type MemoryAccountsRepository struct {
	sync.Mutex
	Accounts map[string]*entity.Account // Map to store accounts, using email as the key
}

func NewMemoryAccountsRepository() *MemoryAccountsRepository {
	return &MemoryAccountsRepository{
		Accounts: make(map[string]*entity.Account),
	}
}

func (repo *MemoryAccountsRepository) CreateAccount(ctx context.Context, account *entity.Account) error {
	repo.Lock()
	defer repo.Unlock()

	if _, exists := repo.Accounts[account.Email]; exists {
		err := appError.NewErrorEmailAlreadyTaken()
		return err
	}

	repo.Accounts[account.Email] = account
	return nil
}

func (repo *MemoryAccountsRepository) UpdateAccount(ctx context.Context, email string, updatedAccount *entity.Account) error {
	repo.Lock()
	defer repo.Unlock()

	existingAccount, exists := repo.Accounts[email]
	if !exists {
		return appError.NewErrorAccountNotFound(email)
	}

	if updatedAccount.Name != "" {
		existingAccount.Name = updatedAccount.Name
	}
	if updatedAccount.AvatarUrl != "" {
		existingAccount.AvatarUrl = updatedAccount.AvatarUrl
	}
	if updatedAccount.Email != "" {
		existingAccount.Email = updatedAccount.Email
	}
	if updatedAccount.IsAccountVerified {
		existingAccount.IsAccountVerified = updatedAccount.IsAccountVerified
	}

	repo.Accounts[email] = existingAccount
	return nil
}

func (repo *MemoryAccountsRepository) GetAccountByID(ctx context.Context, accountID string) (*entity.Account, error) {
	repo.Lock()
	defer repo.Unlock()

	for _, account := range repo.Accounts {
		if account.ID.String() == accountID {
			return account, nil
		}
	}

	return nil, appError.NewErrorAccountNotFound(accountID)
}

func (repo *MemoryAccountsRepository) FindAccountByEmail(ctx context.Context, email string) (*entity.Account, error) {
	repo.Lock()
	defer repo.Unlock()

	account, exists := repo.Accounts[email]

	if !exists {
		return nil, appError.NewErrorAccountNotFound(email)
	}

	return account, nil
}
