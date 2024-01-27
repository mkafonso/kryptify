package factories_test

import (
	repository "kryptify/repositories/memory-repositories"
	"kryptify/usecase"
)

func MakeUpdateCredentialUseCase(
	accountRepo *repository.MemoryAccountsRepository,
	credentialRepo *repository.MemoryCredentialsRepository,
) *usecase.UpdateCredential {
	usecase := usecase.NewUpdateCredential(accountRepo, credentialRepo)
	return usecase
}
