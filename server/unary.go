package main

import (
	"context"

	pb "github.com/herulobarto/grpc-demo-yt/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HeloResponse, error) {
	return &pb.HeloResponse{
		Message: "Hello ",
	}, nil
}
