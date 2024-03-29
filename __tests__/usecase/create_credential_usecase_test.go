package usecase_test

import (
	"context"
	factory_test "kryptify/__tests__/factory"
	memory_repository "kryptify/repository/memory-repository"
	"kryptify/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCredentialUseCase_ShouldCreateNewCredential(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	credentialRepo := memory_repository.NewMemoryCredentialsRepository()
	uc := factory_test.MakeCreateCredentialUseCase(accountRepo, credentialRepo)

	// create an account
	account := factory_test.MakeAccount() // jane@email.com
	accountRepo.CreateAccount(context.Background(), account)

	request := &usecase.CreateCredentialRequest{
		Email:                "jane.doe@email.com",
		Password:             "test1234",
		Website:              "https://my-website.com",
		RequestedByAccountID: account.ID.String(),
	}

	response, err := uc.Execute(context.Background(), request)

	assert.NotNil(t, response)
	assert.NoError(t, err)
	assert.Equal(t, request.Email, response.Credential.Email)
	assert.Equal(t, request.Website, response.Credential.Website)
	assert.Equal(t, request.RequestedByAccountID, response.Credential.OwnerID)
}

func TestCreateCredentialUseCase_TestAccountNotFound(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	credentialRepo := memory_repository.NewMemoryCredentialsRepository()
	uc := factory_test.MakeCreateCredentialUseCase(accountRepo, credentialRepo)

	request := &usecase.CreateCredentialRequest{
		Email:                "jane.doe@email.com",
		Password:             "test1234",
		Website:              "https://my-website.com",
		RequestedByAccountID: "c92fdcdb-8e4b-4b0a-865c-bbc646a467a5", // wrong account
	}

	response, err := uc.Execute(context.Background(), request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, err.Error(), "account not found for: c92fdcdb-8e4b-4b0a-865c-bbc646a467a5")
}
