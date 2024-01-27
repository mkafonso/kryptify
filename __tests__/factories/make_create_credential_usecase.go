package factories_test

import (
	repository "kryptify/repositories/memory-repositories"
	"kryptify/usecase"
)

func MakeCreateCredentialUseCase(
	accountRepo *repository.MemoryAccountsRepository,
	credentialRepo *repository.MemoryCredentialsRepository,
) *usecase.CreateCredential {
	usecase := usecase.NewCreateCredential(accountRepo, credentialRepo)
	return usecase
}
