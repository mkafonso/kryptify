package factory_test

import (
	repository "kryptify/repository/memory-repository"
	"kryptify/usecase"
)

func MakeUpdateAccountUseCase(accountRepo *repository.MemoryAccountsRepository) *usecase.UpdateAccount {
	usecase := usecase.NewUpdateAccount(accountRepo)
	return usecase
}
