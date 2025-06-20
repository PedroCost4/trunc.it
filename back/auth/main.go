package main

import (
	"Trunc-it/trunc.it/auth/config"
	pb "Trunc-it/trunc.it/auth/generated"
	"Trunc-it/trunc.it/auth/handlers"
	"context"
	"flag"
	"log"
)

var (
	port = flag.Int("port", 3000, "The server port")
)

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
}

func (s *AuthServiceServer) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	return handlers.SignUp(ctx, req)
}

func (s *AuthServiceServer) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	return handlers.SignIn(ctx, req)
}

func (s *AuthServiceServer) SignOut(ctx context.Context, req *pb.SignOutRequest) (*pb.SignOutResponse, error) {
	return handlers.SignOut(ctx, req)
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
