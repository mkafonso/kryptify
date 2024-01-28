package main

import (
	"database/sql"
	"kryptify/gapi"
	"kryptify/pb"
	store "kryptify/repository/postgres-repository"
	"kryptify/util"
	"log"
	"net"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadEnvironmentVariable(".")
	if err != nil {
		log.Fatal("cannot load environment variables", err)
	}

	// connect to db
	db, err := sql.Open(config.DB_DRIVER, config.DB_SOURCE)
	if err != nil {
		panic(err)
	}

	store := store.NewPostgresRepository(db)
	initGRPCServer(config, store)
}

func initGRPCServer(config util.Config, store *store.PostgresRepository) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterKryptifyServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPC_SERVER_ADDRESS)
	if err != nil {
		log.Fatal("cannot create listener: ", err)
	}

	log.Printf("Start gRPC server at %s ", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server: ", err)
	}
}
