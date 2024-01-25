package factories_test

import (
	repository "kryptify/repositories/memory-repositories"
	"kryptify/usecases"
)

func MakeCreateCredentialUseCase(credentialRepo *repository.MemoryCredentialsRepository) *usecases.CreateCredential {
	usecase := usecases.NewCreateCredential(credentialRepo)
	return usecase
}
