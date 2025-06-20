package helper

import (
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "Trunc-it/trunc.it/shortener/generated"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func TesterClient() (pb.ShortenerServiceClient, error) {
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()
	c := pb.NewShortenerServiceClient(conn)
	return c, nil
}
