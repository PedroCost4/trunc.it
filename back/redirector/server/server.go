package server

import (
	"context"

	pb "trunc-it/trunc.it/redirector/generated"
	handlers "trunc-it/trunc.it/redirector/handlers"
)

type RedirectorServiceServer struct {
	pb.UnimplementedRedirectorServiceServer
}

func (s *RedirectorServiceServer) GetUrl(ctx context.Context, req *pb.GetUrlRequest) (*pb.GetUrlResponse, error) {
	return handlers.GetUrl(ctx, req)
}
