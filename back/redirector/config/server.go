package config

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"

	pb "trunc-it/trunc.it/redirector/generated"

	"github.com/joho/godotenv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func SetupServer(port int, service pb.RedirectorServiceServer) (*grpc.Server, net.Listener, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		return nil, nil, fmt.Errorf("failed to listen: %w", err)
	}

	s := grpc.NewServer()
	pb.RegisterRedirectorServiceServer(s, service)

	healthServer := health.NewServer()
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(s, healthServer)

	return s, lis, nil
}

func SetupTestClient() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}
	return conn, nil
}

func SetupTestServerAndClient(t *testing.T, service pb.RedirectorServiceServer) (*grpc.ClientConn, func()) {
	err := godotenv.Load(Dir("env.test"))

	if err != nil {
		t.Fatalf("Failed to load env variables: %v", err)
	}

	server, lis, err := SetupServer(3000, service)

	if err != nil {
		t.Fatalf("Failed to setup test server: %v", err)
	}

	go server.Serve(lis)

	conn, err := SetupTestClient()

	if err != nil {
		t.Fatalf("Failed to setup test client: %v", err)
	}

	err = SetupRedis()

	if err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
	}

	_, err = SetupDb()

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return conn, func() {
		redis := GetRedis()
		redis.FlushAll(context.Background())

		db := GetDb()

		_, err := db.sqlxDB.Exec(`
      BEGIN;

      SET session_replication_role = replica;

      DO
      $$
      DECLARE
        tbl RECORD;
      BEGIN
        FOR tbl IN
          SELECT table_schema, table_name
          FROM   information_schema.tables
          WHERE  table_schema = 'public'
            AND  table_type   = 'BASE TABLE'
        LOOP
          EXECUTE format(
            'DROP TABLE IF EXISTS %I.%I;',
            tbl.table_schema,
            tbl.table_name
          );
        END LOOP;
      END
      $$;

      SET session_replication_role = DEFAULT;

      COMMIT;
		  `)

		if err != nil {
			log.Panicf("error: %v", err)
		}

		server.Stop()
		conn.Close()
	}
}
