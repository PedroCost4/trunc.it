package main

import (
	"context"
	"testing"
	"time"

	"Trunc-it/trunc.it/auth/config"
	pb "Trunc-it/trunc.it/auth/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// Dummy service implementation for tests
type testAuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
}

func TestHealthCheck(t *testing.T) {
	srv, lis, err := config.SetupServer(3000, &testAuthServiceServer{})
	if err != nil {
		t.Fatalf("failed to set up server: %v", err)
	}

	go srv.Serve(lis)
	defer srv.Stop()

	time.Sleep(100 * time.Millisecond)

	conn, err := grpc.NewClient("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

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
