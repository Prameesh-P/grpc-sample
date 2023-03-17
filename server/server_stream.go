package main

import (
	"log"
	"time"

	pb "github.com/Prameesh-P/grpc-sample/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GrpcService_SayHelloServerSideStreamingServer)error {
	log.Printf("Got request with names : %v ",req.Names)
	for _,name:=range req.Names{
		res:=&pb.HelloResponse{
			Message:"Hello "+ name,
		}
		if err := stream.Send(res);err!=nil{
			return err
		}
		time.Sleep(2* time.Second)
	}
	return nil
}