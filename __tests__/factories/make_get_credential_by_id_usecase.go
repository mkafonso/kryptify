package factories_test

import (
	repository "kryptify/repositories/memory-repositories"
	"kryptify/usecases"
)

func MakeGetCredentialByIDUseCase(
	accountRepo *repository.MemoryAccountsRepository,
	credentialRepo *repository.MemoryCredentialsRepository,
) *usecases.GetCredentialByID {
	usecase := usecases.NewGetCredentialByID(accountRepo, credentialRepo)
	return usecase
}
