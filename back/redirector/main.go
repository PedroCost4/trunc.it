package main

import (
	"context"
	"flag"
	"log"
	"trunc-it/trunc.it/redirector/config"
	pb "trunc-it/trunc.it/redirector/generated"
	handlers "trunc-it/trunc.it/redirector/handlers"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type RedirectorServiceServer struct {
	pb.UnimplementedRedirectorServiceServer
}

func (s *RedirectorServiceServer) GetUrl(ctx context.Context, req *pb.GetUrlRequest) (*pb.GetUrlResponse, error) {
	return handlers.GetUrl(ctx, req)
}

func main() {
	s, lis, err := config.SetupServer(*port, &RedirectorServiceServer{})

	if err != nil {
		log.Fatalf("Failed to setup server %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
