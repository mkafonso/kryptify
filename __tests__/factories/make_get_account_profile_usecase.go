package factories_test

import (
	repository "kryptify/repository/memory-repository"
	"kryptify/usecase"
)

func MakeGetAccountProfileUseCase(accountRepo *repository.MemoryAccountsRepository) *usecase.GetAccountProfile {
	usecase := usecase.NewGetAccountProfile(accountRepo)
	return usecase
}
