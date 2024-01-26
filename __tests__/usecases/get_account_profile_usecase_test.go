package usecases_test

import (
	"context"
	factories_test "kryptify/__tests__/factories"
	memory_repository "kryptify/repositories/memory-repositories"
	"kryptify/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAccountProfileUseCase_ShouldGetTheProfile(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	usecase := factories_test.MakeGetAccountProfileUseCase(accountRepo)

	// create an account
	account := factories_test.MakeAccount() // jane@email.com
	accountRepo.CreateAccount(context.Background(), account)

	request := &usecases.GetAccountProfileRequest{
		TargetAccountEmail:   account.Email,
		RequestedByAccountID: account.ID.String(),
	}

	response, err := usecase.Execute(context.Background(), request)

	assert.NotNil(t, response)
	assert.NoError(t, err)
	assert.Equal(t, "jane@email.com", response.Account.Email)
}

func TestGetAccountProfileUseCase_TestAccountNotFound(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	usecase := factories_test.MakeGetAccountProfileUseCase(accountRepo)

	// create an account
	account := factories_test.MakeAccount() // jane@email.com
	accountRepo.CreateAccount(context.Background(), account)

	request := &usecases.GetAccountProfileRequest{
		TargetAccountEmail:   account.Email,
		RequestedByAccountID: "c92fdcdb-8e4b-4b0a-865c-bbc646a467a5", // wrong account
	}

	response, err := usecase.Execute(context.Background(), request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, err.Error(), "account not found for: c92fdcdb-8e4b-4b0a-865c-bbc646a467a5")
}

func TestGetAccountProfileUseCase_TestMissingPermission(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	usecase := factories_test.MakeGetAccountProfileUseCase(accountRepo)

	// create an account
	account := factories_test.MakeAccount() // jane@email.com
	accountRepo.CreateAccount(context.Background(), account)

	request := &usecases.GetAccountProfileRequest{
		TargetAccountEmail:   "email-without-permission@email.com",
		RequestedByAccountID: account.ID.String(),
	}

	response, err := usecase.Execute(context.Background(), request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, err.Error(), "missing permission")
}
