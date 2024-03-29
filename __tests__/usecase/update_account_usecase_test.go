package usecase_test

import (
	"context"
	factory_test "kryptify/__tests__/factory"
	memory_repository "kryptify/repository/memory-repository"
	"kryptify/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateAccountUseCase_ShouldUpdateAccount(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	uc := factory_test.MakeUpdateAccountUseCase(accountRepo)

	// create an account
	account := factory_test.MakeAccount() // jane@email.com
	accountRepo.CreateAccount(context.Background(), account)

	assert.Equal(t, account.Name, "Jane Doe")

	request := &usecase.UpdateAccountRequest{
		TargetAccountID:      account.ID.String(),
		RequestedByAccountID: account.ID.String(),
		Name:                 "Jane Smith Doe",
	}

	response, err := uc.Execute(context.Background(), request)

	assert.NotNil(t, response)
	assert.NoError(t, err)
	assert.Equal(t, account.Name, "Jane Smith Doe")
}

func TestUpdateAccountUseCase_CheckUpdatedAtDate(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	uc := factory_test.MakeUpdateAccountUseCase(accountRepo)

	// create an account
	account := factory_test.MakeAccount() // jane@email.com
	accountRepo.CreateAccount(context.Background(), account)

	assert.Equal(t, account.Name, "Jane Doe")

	request := &usecase.UpdateAccountRequest{
		TargetAccountID:      account.ID.String(),
		RequestedByAccountID: account.ID.String(),
		Name:                 "Jane Smith Doe",
	}

	response, err := uc.Execute(context.Background(), request)
	assert.NotNil(t, response)
	assert.NoError(t, err)
	assert.True(t, response.Account.UpdatedAt.After(account.CreatedAt))
}

func TestUpdateAccountUseCase_TestAccountNotFound(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	uc := factory_test.MakeUpdateAccountUseCase(accountRepo)

	// create an account
	account := factory_test.MakeAccount() // jane@email.com
	accountRepo.CreateAccount(context.Background(), account)

	request := &usecase.UpdateAccountRequest{
		TargetAccountID:      account.ID.String(),
		RequestedByAccountID: "c92fdcdb-8e4b-4b0a-865c-bbc646a467a5", // wrong account
		Name:                 "Jane Smith Doe",
	}

	response, err := uc.Execute(context.Background(), request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, err.Error(), "account not found for: c92fdcdb-8e4b-4b0a-865c-bbc646a467a5")
}

func TestUpdateAccountUseCase_TestMissingPermission(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	uc := factory_test.MakeUpdateAccountUseCase(accountRepo)

	// create an account
	account := factory_test.MakeAccount() // jane@email.com
	accountRepo.CreateAccount(context.Background(), account)

	request := &usecase.UpdateAccountRequest{
		TargetAccountID:      "c92fdcdb-8e4b-4b0a-865c-bbc646a467a5",
		RequestedByAccountID: account.ID.String(),
		Name:                 "Jane Smith Doe",
	}

	response, err := uc.Execute(context.Background(), request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, err.Error(), "missing permission")
}
