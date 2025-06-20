package main

import (
	"Trunc-it/trunc.it/shortener/config"
	pb "Trunc-it/trunc.it/shortener/generated"
	"Trunc-it/trunc.it/shortener/handlers"
	"context"
	"flag"
	"log"
)

var (
	port = flag.Int("port", 3000, "The server port")
)

type AuthServiceServer struct {
	pb.UnimplementedShortenerServiceServer
}

func (s *AuthServiceServer) Shorten(ctx context.Context, req *pb.ShortenRequest) (*pb.ShortenResponse, error) {
	return handlers.Shorten(ctx, req)
}

func main() {
	s, lis, err := config.SetupServer(*port, &AuthServiceServer{})
	if err != nil {
		log.Fatalf("Failed to setup server %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
