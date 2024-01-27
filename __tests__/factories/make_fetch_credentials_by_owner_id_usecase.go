package factories_test

import (
	repository "kryptify/repositories/memory-repositories"
	"kryptify/usecase"
)

func MakeFetchCredentialsByOwnerIDUseCase(
	accountRepo *repository.MemoryAccountsRepository,
	credentialRepo *repository.MemoryCredentialsRepository,
) *usecase.FetchCredentialsByOwnerID {
	usecase := usecase.NewFetchCredentialsByOwnerID(accountRepo, credentialRepo)
	return usecase
}
