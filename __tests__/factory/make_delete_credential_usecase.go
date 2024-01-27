package factory_test

import (
	repository "kryptify/repository/memory-repository"
	"kryptify/usecase"
)

func MakeDeleteCredentialUseCase(
	accountRepo *repository.MemoryAccountsRepository,
	credentialRepo *repository.MemoryCredentialsRepository,
) *usecase.DeleteCredential {
	usecase := usecase.NewDeleteCredential(accountRepo, credentialRepo)
	return usecase
}
