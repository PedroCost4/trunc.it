package main

import (
	"context"
	"testing"
	"time"

	"trunc-it/trunc.it/redirector/config"
	pb "trunc-it/trunc.it/redirector/generated"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type testRedirectorServiceServer struct {
	pb.UnimplementedRedirectorServiceServer
}

func TestHealthCheck(t *testing.T) {
	conn, cleanup := config.SetupTestServerAndClient(t, testRedirectorServiceServer{})

	defer cleanup()

	healthClient := healthpb.NewHealthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := healthClient.Check(ctx, &healthpb.HealthCheckRequest{
		Service: "",
	})
	if err != nil {
		t.Fatalf("health check failed: %v", err)
	}

	if resp.GetStatus() != healthpb.HealthCheckResponse_SERVING {
		t.Errorf("expected SERVING, got %v", resp.GetStatus())
	}
}
