package factories_test

import (
	repository "kryptify/repositories/memory-repositories"
	"kryptify/usecases"
)

func MakeGetAccountProfileUseCase(accountRepo *repository.MemoryAccountsRepository) *usecases.GetAccountProfile {
	usecase := usecases.NewGetAccountProfile(accountRepo)
	return usecase
}
