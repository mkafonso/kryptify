package factories_test

import (
	repository "kryptify/repositories/memory-repositories"
	"kryptify/usecase"
)

func MakeGetCredentialByIDUseCase(
	accountRepo *repository.MemoryAccountsRepository,
	credentialRepo *repository.MemoryCredentialsRepository,
) *usecase.GetCredentialByID {
	usecase := usecase.NewGetCredentialByID(accountRepo, credentialRepo)
	return usecase
}
