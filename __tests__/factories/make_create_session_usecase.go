package factories_test

import (
	repository "kryptify/repositories/memory-repositories"
	"kryptify/usecase"
)

func MakeCreateSessionUseCase(
	accountRepo *repository.MemoryAccountsRepository,
	sessionRepo *repository.MemorySessionsRepository,
) *usecase.CreateSession {
	usecase := usecase.NewCreateSession(accountRepo, sessionRepo)
	return usecase
}
