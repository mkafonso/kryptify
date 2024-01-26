package factories_test

import (
	repository "kryptify/repositories/memory-repositories"
	"kryptify/usecases"
)

func MakeCreateAccountUseCase(accountRepo *repository.MemoryAccountsRepository) *usecases.CreateAccount {
	usecase := usecases.NewCreateAccount(accountRepo)
	return usecase
}
