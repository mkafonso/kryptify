package usecase_test

import (
	"context"
	factory_test "kryptify/__tests__/factory"
	memory_repository "kryptify/repository/memory-repository"
	"kryptify/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchCredentialsByOwnerIDUseCase_ShoulFetchCredentialsByOwnerID(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	credentialRepo := memory_repository.NewMemoryCredentialsRepository()
	uc := factory_test.MakeFetchCredentialsByOwnerIDUseCase(accountRepo, credentialRepo)

	// create an account
	account := factory_test.MakeAccount()
	accountRepo.CreateAccount(context.Background(), account)

	// create a credential
	credential := factory_test.MakeCredential("", "", "", account.ID.String()) //make sure it's created by the same account ID
	credentialRepo.CreateCredential(context.Background(), credential)

	request := &usecase.FetchCredentialsByOwnerIDRequest{
		RequestedByAccountID: account.ID.String(),
	}

	response, err := uc.Execute(context.Background(), request)

	assert.NotNil(t, response)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(response.Credentials))
}

func TestFetchCredentialsByOwnerIDUseCase_TestAccountNotFound(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	credentialRepo := memory_repository.NewMemoryCredentialsRepository()
	uc := factory_test.MakeFetchCredentialsByOwnerIDUseCase(accountRepo, credentialRepo)

	// create an account
	account := factory_test.MakeAccount()
	accountRepo.CreateAccount(context.Background(), account)

	// create a credential
	credential := factory_test.MakeCredential() // will be created by random accountID
	credentialRepo.CreateCredential(context.Background(), credential)

	request := &usecase.FetchCredentialsByOwnerIDRequest{
		RequestedByAccountID: credential.OwnerID,
	}

	response, err := uc.Execute(context.Background(), request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, err.Error(), "account not found for: c92fdcdb-8e4b-4b0a-865c-bbc646a467a5")
}
