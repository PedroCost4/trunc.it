package config

import (
	"fmt"
	"net"

	pb "Trunc-it/trunc.it/shortener/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func SetupServer(port int, service pb.ShortenerServiceServer) (*grpc.Server, net.Listener, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to listen: %w", err)
	}

	s := grpc.NewServer()
	pb.RegisterShortenerServiceServer(s, service)

	healthServer := health.NewServer()
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(s, healthServer)

	return s, lis, nil
}
