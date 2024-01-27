package usecase_test

import (
	"context"
	factory_test "kryptify/__tests__/factory"
	memory_repository "kryptify/repository/memory-repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateCredentialUseCase_ShouldUpdateCredential(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	credentialRepo := memory_repository.NewMemoryCredentialsRepository()
	usecase := factory_test.MakeUpdateCredentialUseCase(accountRepo, credentialRepo)

	// create an account
	account := factory_test.MakeAccount()
	accountRepo.CreateAccount(context.Background(), account)

	// create a credential
	credential := factory_test.MakeCredential("", "", "", account.ID.String()) //make sure it's created by the same account ID
	credentialRepo.CreateCredential(context.Background(), credential)

	assert.Equal(t, "", credential.Category)

	request := &usecase.UpdateCredentialRequest{
		TargetCredentialID:   credential.ID.String(),
		RequestedByAccountID: account.ID.String(),
		Category:             "Social Network",
	}

	response, err := usecase.Execute(context.Background(), request)

	assert.NotNil(t, response)
	assert.NoError(t, err)
	assert.Equal(t, "Social Network", response.Credential.Category)
}

func TestUpdateCredentialUseCase_TestAccountNotFound(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	credentialRepo := memory_repository.NewMemoryCredentialsRepository()
	usecase := factory_test.MakeUpdateCredentialUseCase(accountRepo, credentialRepo)

	// create an account
	account := factory_test.MakeAccount()
	accountRepo.CreateAccount(context.Background(), account)

	// create a credential
	credential := factory_test.MakeCredential() // will be created by random accountID
	credentialRepo.CreateCredential(context.Background(), credential)

	assert.Equal(t, "", credential.Category)

	request := &usecase.UpdateCredentialRequest{
		TargetCredentialID:   credential.ID.String(),
		RequestedByAccountID: credential.OwnerID,
		Category:             "Social Network",
	}

	response, err := usecase.Execute(context.Background(), request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, err.Error(), "account not found for: c92fdcdb-8e4b-4b0a-865c-bbc646a467a5")
}

func TestUpdateCredentialUseCase_TestMissingPermission(t *testing.T) {
	accountRepo := memory_repository.NewMemoryAccountsRepository()
	credentialRepo := memory_repository.NewMemoryCredentialsRepository()
	usecase := factory_test.MakeUpdateCredentialUseCase(accountRepo, credentialRepo)

	// create an account
	account := factory_test.MakeAccount()
	accountRepo.CreateAccount(context.Background(), account)

	// create a credential
	credential := factory_test.MakeCredential() // will be created by random accountID
	credentialRepo.CreateCredential(context.Background(), credential)

	assert.Equal(t, "", credential.Category)

	request := &usecase.UpdateCredentialRequest{
		TargetCredentialID:   credential.ID.String(),
		RequestedByAccountID: account.ID.String(),
	}

	response, err := usecase.Execute(context.Background(), request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, err.Error(), "missing permission")
}
