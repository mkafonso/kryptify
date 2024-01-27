package factories_test

import (
	repository "kryptify/repository/memory-repository"
	"kryptify/usecase"
)

func MakeCreateCredentialUseCase(
	accountRepo *repository.MemoryAccountsRepository,
	credentialRepo *repository.MemoryCredentialsRepository,
) *usecase.CreateCredential {
	usecase := usecase.NewCreateCredential(accountRepo, credentialRepo)
	return usecase
}
