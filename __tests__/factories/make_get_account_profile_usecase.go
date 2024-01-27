package factories_test

import (
	repository "kryptify/repositories/memory-repositories"
	"kryptify/usecase"
)

func MakeGetAccountProfileUseCase(accountRepo *repository.MemoryAccountsRepository) *usecase.GetAccountProfile {
	usecase := usecase.NewGetAccountProfile(accountRepo)
	return usecase
}
