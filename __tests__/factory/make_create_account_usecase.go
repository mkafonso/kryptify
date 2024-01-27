package factory_test

import (
	repository "kryptify/repository/memory-repository"
	"kryptify/usecase"
)

func MakeCreateAccountUseCase(accountRepo *repository.MemoryAccountsRepository) *usecase.CreateAccount {
	usecase := usecase.NewCreateAccount(accountRepo)
	return usecase
}
