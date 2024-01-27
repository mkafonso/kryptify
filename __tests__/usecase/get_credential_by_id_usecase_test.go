package usecase_test

import (
	"context"
	factory_test "kryptify/__tests__/factory"
	memory_repository "kryptify/repository/memory-repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCredentialByIDUseCase_ShouldGetCredentialByID(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	credentialRepo := memory_repository.NewMemoryCredentialsRepository()
	usecase := factory_test.MakeGetCredentialByIDUseCase(accountRepo, credentialRepo)

	// create an account
	account := factory_test.MakeAccount()
	accountRepo.CreateAccount(context.Background(), account)

	// create a credential
	credential := factory_test.MakeCredential("", "", "", account.ID.String()) //make sure it's created by the same account ID
	credentialRepo.CreateCredential(context.Background(), credential)

	request := &usecase.GetCredentialByIDRequest{
		TargetCredentialID:   credential.ID.String(),
		RequestedByAccountID: account.ID.String(),
	}

	response, err := usecase.Execute(context.Background(), request)

	assert.NotNil(t, response)
	assert.NoError(t, err)
	assert.Equal(t, "https://my-website.com", response.Credential.Website)
	assert.Equal(t, "john@email.com", response.Credential.Email)
}

func TestGetCredentialByIDUseCase_TestAccountNotFound(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	credentialRepo := memory_repository.NewMemoryCredentialsRepository()
	usecase := factory_test.MakeGetCredentialByIDUseCase(accountRepo, credentialRepo)

	// create an account
	account := factory_test.MakeAccount()
	accountRepo.CreateAccount(context.Background(), account)

	// create a credential
	credential := factory_test.MakeCredential() // will be created by random accountID
	credentialRepo.CreateCredential(context.Background(), credential)

	request := &usecase.GetCredentialByIDRequest{
		TargetCredentialID:   credential.ID.String(),
		RequestedByAccountID: credential.OwnerID,
	}

	response, err := usecase.Execute(context.Background(), request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, err.Error(), "account not found for: c92fdcdb-8e4b-4b0a-865c-bbc646a467a5")
}

func TestGetCredentialByIDUseCase_TestMissingPermission(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	credentialRepo := memory_repository.NewMemoryCredentialsRepository()
	usecase := factory_test.MakeGetCredentialByIDUseCase(accountRepo, credentialRepo)

	// create an account
	account := factory_test.MakeAccount()
	accountRepo.CreateAccount(context.Background(), account)

	// create a credential
	credential := factory_test.MakeCredential() // will be created by random accountID
	credentialRepo.CreateCredential(context.Background(), credential)

	request := &usecase.GetCredentialByIDRequest{
		TargetCredentialID:   credential.ID.String(),
		RequestedByAccountID: account.ID.String(),
	}

	response, err := usecase.Execute(context.Background(), request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, err.Error(), "missing permission")
}
