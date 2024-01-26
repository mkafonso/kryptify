package usecases_test

import (
	"context"
	factories_test "kryptify/__tests__/factories"
	memory_repository "kryptify/repositories/memory-repositories"
	usecases "kryptify/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteCredentialUseCase_ShouldCreateNewCredential(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	credentialRepo := memory_repository.NewMemoryCredentialsRepository()
	usecase := factories_test.MakeDeleteCredentialUseCase(accountRepo, credentialRepo)

	// create an account
	account := factories_test.MakeAccount()
	accountRepo.CreateAccount(context.Background(), account)

	// create a credential
	credential := factories_test.MakeCredential("", "", "", account.ID.String()) //make sure it's created by the same account ID
	credentialRepo.CreateCredential(context.Background(), credential)

	request := &usecases.DeleteCredentialRequest{
		TargetCredentialID:   credential.ID.String(),
		RequestedByAccountID: account.ID.String(),
	}

	response, err := usecase.Execute(context.Background(), request)

	assert.NotNil(t, response)
	assert.NoError(t, err)
}

func TestDeleteCredentialUseCase_TestAccountNotFound(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	credentialRepo := memory_repository.NewMemoryCredentialsRepository()
	usecase := factories_test.MakeDeleteCredentialUseCase(accountRepo, credentialRepo)

	// create an account
	account := factories_test.MakeAccount()
	accountRepo.CreateAccount(context.Background(), account)

	// create a credential
	credential := factories_test.MakeCredential() // will be created by random accountID
	credentialRepo.CreateCredential(context.Background(), credential)

	request := &usecases.DeleteCredentialRequest{
		TargetCredentialID:   credential.ID.String(),
		RequestedByAccountID: credential.OwnerID,
	}

	response, err := usecase.Execute(context.Background(), request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, err.Error(), "account not found for: c92fdcdb-8e4b-4b0a-865c-bbc646a467a5")
}

func TestDeleteCredentialUseCase_TestMissingPermission(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	credentialRepo := memory_repository.NewMemoryCredentialsRepository()
	usecase := factories_test.MakeDeleteCredentialUseCase(accountRepo, credentialRepo)

	// create an account
	account := factories_test.MakeAccount()
	accountRepo.CreateAccount(context.Background(), account)

	// create a credential
	credential := factories_test.MakeCredential() // will be created by random accountID
	credentialRepo.CreateCredential(context.Background(), credential)

	request := &usecases.DeleteCredentialRequest{
		TargetCredentialID:   credential.ID.String(),
		RequestedByAccountID: account.ID.String(),
	}

	response, err := usecase.Execute(context.Background(), request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, err.Error(), "missing permission")
}
