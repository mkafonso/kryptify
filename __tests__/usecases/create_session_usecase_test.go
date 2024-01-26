package usecases_test

import (
	"context"
	factories_test "kryptify/__tests__/factories"
	memory_repository "kryptify/repositories/memory-repositories"
	usecases "kryptify/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSessionUseCase_ShouldCreateSession(t *testing.T) {
	accountsRepo := memory_repository.NewMemoryAccountsRepository()
	sessionsRepo := memory_repository.NewMemorySessionsRepository()
	usecase := factories_test.MakeCreateSessionUseCase(accountsRepo, sessionsRepo)

	// create an account
	account := factories_test.MakeAccount() // jane@email.com
	accountsRepo.CreateAccount(context.Background(), account)

	request := &usecases.CreateSessionRequest{
		Email:     "jane@email.com",
		Password:  "myVerySecurePassword",
		UserAgent: "valid-user-agent",
		ClientIP:  "valid-client-ip",
	}

	response, err := usecase.Execute(context.Background(), request)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.AccessToken)
	assert.NotEmpty(t, response.AccessTokenExpiresAt)
	assert.NotEmpty(t, response.RefreshToken)
	assert.NotEmpty(t, response.RefreshTokenExpiresAt)
	assert.NotEmpty(t, response.SessionID)
	assert.Equal(t, account.Email, response.Account.Email)
	assert.Equal(t, account.Name, response.Account.Name)
}

func TestCreateSessionUseCase_TestWrongEmailAddress(t *testing.T) {
	accountsRepo := memory_repository.NewMemoryAccountsRepository()
	sessionsRepo := memory_repository.NewMemorySessionsRepository()
	usecase := factories_test.MakeCreateSessionUseCase(accountsRepo, sessionsRepo)

	// create an account
	account := factories_test.MakeAccount() // jane@email.com
	accountsRepo.CreateAccount(context.Background(), account)

	request := &usecases.CreateSessionRequest{
		Email:     "wrong-account@email.com",
		Password:  "myVerySecurePassword",
		UserAgent: "valid-user-agent",
		ClientIP:  "valid-client-ip",
	}

	response, err := usecase.Execute(context.Background(), request)
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, err.Error(), "invalid credentials")
}

func TestCreateSessionUseCase_TestWrongPassword(t *testing.T) {
	accountsRepo := memory_repository.NewMemoryAccountsRepository()
	sessionsRepo := memory_repository.NewMemorySessionsRepository()
	usecase := factories_test.MakeCreateSessionUseCase(accountsRepo, sessionsRepo)

	// create an account
	account := factories_test.MakeAccount() // jane@email.com
	accountsRepo.CreateAccount(context.Background(), account)

	request := &usecases.CreateSessionRequest{
		Email:     "jane@email.com",
		Password:  "wrong-password",
		UserAgent: "valid-user-agent",
		ClientIP:  "valid-client-ip",
	}

	response, err := usecase.Execute(context.Background(), request)
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, err.Error(), "invalid credentials")
}