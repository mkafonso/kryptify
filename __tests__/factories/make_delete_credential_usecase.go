package factories_test

import (
	repository "kryptify/repositories/memory-repositories"
	"kryptify/usecases"
)

func MakeDeleteCredentialUseCase(
	accountRepo *repository.MemoryAccountsRepository,
	credentialRepo *repository.MemoryCredentialsRepository,
) *usecases.DeleteCredential {
	usecase := usecases.NewDeleteCredential(accountRepo, credentialRepo)
	return usecase
}
