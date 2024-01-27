package factories_test

import (
	repository "kryptify/repository/memory-repository"
	"kryptify/usecase"
)

func MakeUpdateCredentialUseCase(
	accountRepo *repository.MemoryAccountsRepository,
	credentialRepo *repository.MemoryCredentialsRepository,
) *usecase.UpdateCredential {
	usecase := usecase.NewUpdateCredential(accountRepo, credentialRepo)
	return usecase
}
