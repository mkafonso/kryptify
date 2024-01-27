package factories_test

import (
	repository "kryptify/repositories/memory-repositories"
	"kryptify/usecase"
)

func MakeDeleteCredentialUseCase(
	accountRepo *repository.MemoryAccountsRepository,
	credentialRepo *repository.MemoryCredentialsRepository,
) *usecase.DeleteCredential {
	usecase := usecase.NewDeleteCredential(accountRepo, credentialRepo)
	return usecase
}
