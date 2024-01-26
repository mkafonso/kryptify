package usecases_test

import (
	"context"
	factories_test "kryptify/__tests__/factories"
	memory_repository "kryptify/repositories/memory-repositories"
	usecases "kryptify/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccountUseCase_ShouldCreateNewAccount(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	usecase := factories_test.MakeCreateAccountUseCase(accountRepo)

	request := &usecases.CreateAccountRequest{
		Name:     "Jane Doe",
		Email:    "jane.doe@email.com",
		Password: "myVerySecurePassword",
	}

	response, err := usecase.Execute(context.Background(), request)

	assert.NotNil(t, response)
	assert.NoError(t, err)
	assert.Equal(t, request.Name, response.Account.Name)
	assert.Equal(t, request.Email, response.Account.Email)
	assert.Equal(t, false, response.Account.IsAccountVerified)
}

func TestCreateAccountUseCase_TestEmailTaken(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	usecase := factories_test.MakeCreateAccountUseCase(accountRepo)

	request := &usecases.CreateAccountRequest{
		Name:     "Jane Doe",
		Email:    "jane.doe@email.com",
		Password: "myVerySecurePassword",
	}

	usecase.Execute(context.Background(), request) // email already taken
	response, err := usecase.Execute(context.Background(), request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, err.Error(), "email already taken")
}
