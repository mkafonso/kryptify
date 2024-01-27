package gapi

import (
	"fmt"
	"kryptify/pb"
	store "kryptify/repository/postgres-repository"
	"kryptify/token"
	"kryptify/util"
)

type Server struct {
	pb.UnimplementedKryptifyServiceServer
	tokenMaker token.MakerInterface
	store      *store.PostgresRepository
}

func NewServer(config util.Config, store *store.PostgresRepository) (*Server, error) {
	tokenSymmetricKey := config.TOKEN_SYMETRIC_KEY
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
