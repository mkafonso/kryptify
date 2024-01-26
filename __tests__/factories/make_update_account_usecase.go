package factories_test

import (
	repository "kryptify/repositories/memory-repositories"
	"kryptify/usecases"
)

func MakeUpdateAccountUseCase(accountRepo *repository.MemoryAccountsRepository) *usecases.UpdateAccount {
	usecase := usecases.NewUpdateAccount(accountRepo)
	return usecase
}
