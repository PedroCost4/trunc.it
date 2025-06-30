package handlers_test

import (
	"context"
	"fmt"
	"testing"
	"time"
	"trunc-it/trunc.it/redirector/config"
	pb "trunc-it/trunc.it/redirector/generated"
	"trunc-it/trunc.it/redirector/helpers"
	"trunc-it/trunc.it/redirector/server"

	"google.golang.org/protobuf/proto"

	"github.com/stretchr/testify/assert"
)

func TestGetUrlCached(t *testing.T) {
	conn, cleanup := config.SetupTestServerAndClient(t, &server.RedirectorServiceServer{})

	defer cleanup()

	request := pb.GetUrlRequest{ShortCode: "abcdef"}

	stringfied, _ := proto.Marshal(&pb.GetUrlResponse{Success: true, Msg: "", Data: "12345678"})

	helpers.Cache.Store(request.ShortCode, stringfied, time.Millisecond*10000)

	client := pb.NewRedirectorServiceClient(conn)

	res, err := client.GetUrl(context.Background(), &request)

	assert.Equal(t, nil, err, fmt.Sprintf("%v", err))
	assert.Equal(t, true, res.Success, res.Success)
	assert.Equal(t, "12345678", res.Data, res.Data)
}

func TestGetUrl(t *testing.T) {
	conn, cleanup := config.SetupTestServerAndClient(t, &server.RedirectorServiceServer{})

	defer cleanup()

	request := pb.GetUrlRequest{ShortCode: "abcdef"}

	client := pb.NewRedirectorServiceClient(conn)

	res, err := client.GetUrl(context.Background(), &request)

	assert.Equal(t, nil, err, fmt.Sprintf("%v", err))
	assert.Equal(t, true, res.Success, res.Success)
	assert.Equal(t, "12345678", res.Data, res.Data)
}
