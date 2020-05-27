package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/marcsj/ocaptchas/challenge"
	"github.com/marcsj/ocaptchas/controller"
	"github.com/marcsj/ocaptchas/repo"
	"github.com/marcsj/ocaptchas/server"
	"github.com/namsral/flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
	"net/http"
)

var (
	databaseFile = flag.String(
		"database_file", "stored.db", "location for database file")
	grpcPort = flag.String(
		"grpc_port", "50051", "port for grpc connections")
	httpPort = flag.String(
		"http_port", "8080", "port for http")
	apiKey = flag.String(
		"api_key", uuid.New().String(), "key for usage of admin api")
)

func main() {
	flag.Parse()

	errChannel := make(chan error)

	db, err := gorm.Open("sqlite3", *databaseFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sessionRepo, err := repo.NewSessionRepo(db)
	if err != nil {
		log.Fatal(err)
	}
	questionsRepo, err := repo.NewChallengeQuestionsRepo(db)
	if err != nil {
		log.Fatal(err)
	}

	challengeController := controller.NewChallengeController(sessionRepo, questionsRepo)

	challengeServer := server.NewChallengeServer(challengeController)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", *grpcPort))
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	mux := runtime.NewServeMux()

	challenge.RegisterChallengeServer(grpcServer, challengeServer)
	err = challenge.RegisterChallengeHandlerServer(context.Background(), mux, challengeServer)
	if err != nil {
		log.Fatal(err)
	}

	// running gRPC server
	go func() {
		grpclog.Infof("Starting gRPC server. tcp port: %v", *grpcPort)
		errChannel <- grpcServer.Serve(lis)
	}()

	// running http server
	go func() {
		grpclog.Infof("Starting http server. http port: %v", *httpPort)
		grpcGateway := http.Server{
			Addr: fmt.Sprintf(":%v", *httpPort),
			Handler: mux,
		}
		errChannel <- grpcGateway.ListenAndServe()
	}()

	log.Println("started ocaptchas server")
	log.Println("api key:", *apiKey)

	for {
		select {
		case err := <-errChannel:
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
