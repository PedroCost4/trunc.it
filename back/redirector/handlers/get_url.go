package handlers

import (
	"context"
	"fmt"
	pb "trunc-it/trunc.it/redirector/generated"
	"trunc-it/trunc.it/redirector/helpers"

	"google.golang.org/protobuf/proto"
)

func GetUrl(ctx context.Context, req *pb.GetUrlRequest) (*pb.GetUrlResponse, error) {
	cacheResponse, err := helpers.Cache.Lookup(req.ShortCode)

	if err == nil && cacheResponse != nil {
		var response pb.GetUrlResponse

		err := proto.Unmarshal([]byte(*cacheResponse), &response)

		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal cached response: %v", err)
		}

		return &response, nil
	}

	return &pb.GetUrlResponse{Success: true, Data: "asdasad"}, err
}
