package factories_test

import (
	repository "kryptify/repositories/memory-repositories"
	"kryptify/usecases"
)

func MakeCreateSessionUseCase(
	accountRepo *repository.MemoryAccountsRepository,
	sessionRepo *repository.MemorySessionsRepository,
) *usecases.CreateSession {
	usecase := usecases.NewCreateSession(accountRepo, sessionRepo)
	return usecase
}
