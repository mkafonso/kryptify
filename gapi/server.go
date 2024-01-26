package gapi

import (
	"fmt"
	"kryptify/pb"
	store "kryptify/repositories/postgres-repositories"
	"kryptify/token"
	"kryptify/utils"
)

type Server struct {
	pb.UnimplementedKryptifyServiceServer
	tokenMaker token.MakerInterface
	store      *store.PostgresRepository
}

func NewServer(config utils.Config, store *store.PostgresRepository) (*Server, error) {
	tokenSymmetricKey := "12345678123456781234567812345678"
	tokenMaker, err := token.NewPasetoMaker(tokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		tokenMaker: tokenMaker,
		store:      store,
	}

	return server, nil
}
