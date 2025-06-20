package helpers

import (
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "trunc-it/trunc.it/redirector/generated"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func TesterClient() (pb.RedirectorServiceClient, error) {
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}

	defer conn.Close()

	c := pb.NewRedirectorServiceClient(conn)

	return c, nil
}
