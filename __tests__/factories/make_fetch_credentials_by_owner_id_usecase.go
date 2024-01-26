package factories_test

import (
	repository "kryptify/repositories/memory-repositories"
	"kryptify/usecases"
)

func MakeFetchCredentialsByOwnerIDUseCase(
	accountRepo *repository.MemoryAccountsRepository,
	credentialRepo *repository.MemoryCredentialsRepository,
) *usecases.FetchCredentialsByOwnerID {
	usecase := usecases.NewFetchCredentialsByOwnerID(accountRepo, credentialRepo)
	return usecase
}
