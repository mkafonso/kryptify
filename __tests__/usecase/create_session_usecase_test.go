package usecase_test

import (
	"context"
	factory_test "kryptify/__tests__/factory"
	memory_repository "kryptify/repository/memory-repository"
	"kryptify/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSessionUseCase_ShouldCreateSession(t *testing.T) {
	accountsRepo := memory_repository.NewMemoryAccountsRepository()
	sessionsRepo := memory_repository.NewMemorySessionsRepository()
	uc := factory_test.MakeCreateSessionUseCase(accountsRepo, sessionsRepo)

	// create an account
	account := factory_test.MakeAccount() // jane@email.com
	accountsRepo.CreateAccount(context.Background(), account)

	request := &usecase.CreateSessionRequest{
		Email:     "jane@email.com",
		Password:  "myVerySecurePassword",
		UserAgent: "valid-user-agent",
		ClientIP:  "valid-client-ip",
	}

	response, err := uc.Execute(context.Background(), request)
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
	uc := factory_test.MakeCreateSessionUseCase(accountsRepo, sessionsRepo)

	// create an account
	account := factory_test.MakeAccount() // jane@email.com
	accountsRepo.CreateAccount(context.Background(), account)

	request := &usecase.CreateSessionRequest{
		Email:     "wrong-account@email.com",
		Password:  "myVerySecurePassword",
		UserAgent: "valid-user-agent",
		ClientIP:  "valid-client-ip",
	}

	response, err := uc.Execute(context.Background(), request)
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, err.Error(), "invalid credentials")
}

func TestCreateSessionUseCase_TestWrongPassword(t *testing.T) {
	accountsRepo := memory_repository.NewMemoryAccountsRepository()
	sessionsRepo := memory_repository.NewMemorySessionsRepository()
	uc := factory_test.MakeCreateSessionUseCase(accountsRepo, sessionsRepo)

	// create an account
	account := factory_test.MakeAccount() // jane@email.com
	accountsRepo.CreateAccount(context.Background(), account)

	request := &usecase.CreateSessionRequest{
		Email:     "jane@email.com",
		Password:  "wrong-password",
		UserAgent: "valid-user-agent",
		ClientIP:  "valid-client-ip",
	}

	response, err := uc.Execute(context.Background(), request)
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, err.Error(), "invalid credentials")
}
