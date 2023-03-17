package main

import (
	"context"

	pb "github.com/Prameesh-P/grpc-sample/proto"
)

func (s *helloServer) SayHello (ctx context.Context , req *pb.NoParam)(*pb.HelloResponse ,error){

	return &pb.HelloResponse{
		Message: "Hello",
	},nil
}