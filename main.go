package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/uacademy/e_commerce/order_service/clients"
	"github.com/uacademy/e_commerce/order_service/config"
	ecom "github.com/uacademy/e_commerce/order_service/proto-gen/e_commerce"
	srv "github.com/uacademy/e_commerce/order_service/services/order"
	"github.com/uacademy/e_commerce/order_service/storage"
	"github.com/uacademy/e_commerce/order_service/storage/postgres"
)

func main() {
	cfg := config.Load()

	psqlConString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	var stg storage.StorageI
	stg, err := postgres.InitDb(psqlConString)
	if err != nil {
		panic(err)
	}

	println("gRPC server tutorial in Go")

	listener, err := net.Listen("tcp", ":9002")
	if err != nil {
		panic(err)
	}

	client, err := clients.NewGrpcClients(cfg)
	if err != nil{
		log.Fatalf("failed to connect product_service: %v", err)
	}

	s := grpc.NewServer()
	ecom.RegisterOrderServiceServer(s, srv.NewOrderService(client, stg))
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
