package main

import (
	"log"
	"net"
	pb "github.com/Prameesh-P/grpc-sample/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)
type helloServer struct{

	pb.GrpcServiceServer

}

func main() {

	lis, err := net.Listen("tcp",port)

	if err != nil{
		log.Fatalf("failed Start the server..!! Got an error %v",err)
	}

	grpcServer:= grpc.NewServer()
	pb.RegisterGrpcServiceServer(grpcServer,&helloServer{})
	log.Printf("Server is started on %v",lis.Addr())
	if err := grpcServer.Serve(lis);err!=nil{
		log.Fatalf("Failed to start grpc %v",err)
	}

}