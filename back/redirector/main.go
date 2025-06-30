package main

import (
	"flag"
	"log"
	"trunc-it/trunc.it/redirector/config"
	"trunc-it/trunc.it/redirector/server"

	"github.com/joho/godotenv"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Failed to load env variables: %v", err)
	}

	s, lis, err := config.SetupServer(*port, &server.RedirectorServiceServer{})

	if err != nil {
		log.Fatalf("Failed to setup server %v", err)
	}

	err = config.SetupRedis()

	if err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
	}

	_, err = config.SetupDb()

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		config.GetDb().Close()

		log.Fatalf("failed to serve: %v", err)
	}
}
