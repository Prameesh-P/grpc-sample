package main

import (
	"log"
	pb "github.com/Prameesh-P/grpc-sample/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {

	conn, err := grpc.Dial("localhost"+port,grpc.WithTransportCredentials(insecure.NewCredentials()))
	
	if err !=nil{
			log.Fatalf("Can't connect..!!! got error %v ",err)
	}
	
	defer conn.Close()

	client := pb.NewGrpcServiceClient(conn)

	callSayHello(client)
}