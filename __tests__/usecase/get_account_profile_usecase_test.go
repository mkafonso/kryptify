package usecase_test

import (
	"context"
	factory_test "kryptify/__tests__/factory"
	memory_repository "kryptify/repository/memory-repository"
	"kryptify/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAccountProfileUseCase_ShouldGetTheProfile(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	uc := factory_test.MakeGetAccountProfileUseCase(accountRepo)

	// create an account
	account := factory_test.MakeAccount() // jane@email.com
	accountRepo.CreateAccount(context.Background(), account)

	request := &usecase.GetAccountProfileRequest{
		TargetAccountEmail:   account.Email,
		RequestedByAccountID: account.ID.String(),
	}

	response, err := uc.Execute(context.Background(), request)

	assert.NotNil(t, response)
	assert.NoError(t, err)
	assert.Equal(t, "jane@email.com", response.Account.Email)
}

func TestGetAccountProfileUseCase_TestAccountNotFound(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	uc := factory_test.MakeGetAccountProfileUseCase(accountRepo)

	// create an account
	account := factory_test.MakeAccount() // jane@email.com
	accountRepo.CreateAccount(context.Background(), account)

	request := &usecase.GetAccountProfileRequest{
		TargetAccountEmail:   account.Email,
		RequestedByAccountID: "c92fdcdb-8e4b-4b0a-865c-bbc646a467a5", // wrong account
	}

	response, err := uc.Execute(context.Background(), request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, err.Error(), "account not found for: c92fdcdb-8e4b-4b0a-865c-bbc646a467a5")
}

func TestGetAccountProfileUseCase_TestMissingPermission(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	uc := factory_test.MakeGetAccountProfileUseCase(accountRepo)

	// create an account
	account := factory_test.MakeAccount() // jane@email.com
	accountRepo.CreateAccount(context.Background(), account)

	request := &usecase.GetAccountProfileRequest{
		TargetAccountEmail:   "email-without-permission@email.com",
		RequestedByAccountID: account.ID.String(),
	}

	response, err := uc.Execute(context.Background(), request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, err.Error(), "missing permission")
}
