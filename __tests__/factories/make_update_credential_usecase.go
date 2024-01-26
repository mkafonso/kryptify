package factories_test

import (
	repository "kryptify/repositories/memory-repositories"
	"kryptify/usecases"
)

func MakeUpdateCredentialUseCase(
	accountRepo *repository.MemoryAccountsRepository,
	credentialRepo *repository.MemoryCredentialsRepository,
) *usecases.UpdateCredential {
	usecase := usecases.NewUpdateCredential(accountRepo, credentialRepo)
	return usecase
}
