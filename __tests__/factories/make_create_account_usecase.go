package factories_test

import (
	repository "kryptify/repositories/memory-repositories"
	"kryptify/usecase"
)

func MakeCreateAccountUseCase(accountRepo *repository.MemoryAccountsRepository) *usecase.CreateAccount {
	usecase := usecase.NewCreateAccount(accountRepo)
	return usecase
}
