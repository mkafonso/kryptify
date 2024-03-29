package usecase_test

import (
	"context"
	factory_test "kryptify/__tests__/factory"
	memory_repository "kryptify/repository/memory-repository"
	"kryptify/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccountUseCase_ShouldCreateNewAccount(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	uc := factory_test.MakeCreateAccountUseCase(accountRepo)

	request := &usecase.CreateAccountRequest{
		Name:     "Jane Doe",
		Email:    "jane.doe@email.com",
		Password: "myVerySecurePassword",
	}

	response, err := uc.Execute(context.Background(), request)

	assert.NotNil(t, response)
	assert.NoError(t, err)
	assert.Equal(t, request.Name, response.Account.Name)
	assert.Equal(t, request.Email, response.Account.Email)
	assert.Equal(t, false, response.Account.IsAccountVerified)
}

func TestCreateAccountUseCase_TestEmailTaken(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	uc := factory_test.MakeCreateAccountUseCase(accountRepo)

	request := &usecase.CreateAccountRequest{
		Name:     "Jane Doe",
		Email:    "jane.doe@email.com",
		Password: "myVerySecurePassword",
	}

	uc.Execute(context.Background(), request) // email already taken
	response, err := uc.Execute(context.Background(), request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, err.Error(), "email already taken")
}
