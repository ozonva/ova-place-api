package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/ozonva/ova-place-api/internal/api"
	"github.com/ozonva/ova-place-api/internal/repo"
	desc "github.com/ozonva/ova-place-api/pkg/ova-place-api"
	"google.golang.org/grpc"
)

const (
	grpcPort = ":7002"
)

var (
	grpcEndpoint = flag.String("grpc-server-endpoint", "0.0.0.0"+grpcPort, "gRPC server endpoint")
)

func runGrpc() error {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalln(err)
	}

	db, err := sqlx.Connect(os.Getenv("DB_DRIVER"), os.Getenv("DB_STRING"))
	if err != nil {
		log.Fatalln(err)
	}

	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterOvaPlaceApiV1Server(s, api.NewOvaPlaceApi(repo.NewRepo(db)))

	fmt.Printf("Server listening on %s\n", *grpcEndpoint)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func main() {
	flag.Parse()

	err := runGrpc()
	if err != nil {
		log.Fatalf("grpc not started")
	}
}
