package factory_test

import (
	repository "kryptify/repository/memory-repository"
	"kryptify/usecase"
)

func MakeFetchCredentialsByOwnerIDUseCase(
	accountRepo *repository.MemoryAccountsRepository,
	credentialRepo *repository.MemoryCredentialsRepository,
) *usecase.FetchCredentialsByOwnerID {
	usecase := usecase.NewFetchCredentialsByOwnerID(accountRepo, credentialRepo)
	return usecase
}
