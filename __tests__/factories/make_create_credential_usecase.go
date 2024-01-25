package factories_test

import (
	repository "kryptify/repositories/memory-repositories"
	"kryptify/usecases"
)

func MakeCreateCredentialUseCase(
	accountRepo *repository.MemoryAccountsRepository,
	credentialRepo *repository.MemoryCredentialsRepository,
) *usecases.CreateCredential {
	usecase := usecases.NewCreateCredential(accountRepo, credentialRepo)
	return usecase
}
