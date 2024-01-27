package factories_test

import (
	repository "kryptify/repositories/memory-repositories"
	"kryptify/usecase"
)

func MakeUpdateAccountUseCase(accountRepo *repository.MemoryAccountsRepository) *usecase.UpdateAccount {
	usecase := usecase.NewUpdateAccount(accountRepo)
	return usecase
}
