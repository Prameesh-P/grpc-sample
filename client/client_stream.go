package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Prameesh-P/grpc-sample/proto"
)

func callSayHelloClientStream(client pb.GrpcServiceClient,names *pb.NamesList){

	log.Printf("Client stream is started ..!! ")
	stream ,err :=client.SayHelloClientStreaming(context.Background())
	if err !=nil{
		log.Fatalf("Could not send the names %v",err)
	}
	for _,name:=range names.Names{
		req:=&pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req);err!=nil{
			log.Fatalf("Could not send the request got an error %v",err)
		}
		time.Sleep(2*time.Second)
	}
	res,err:=stream.CloseAndRecv()
	log.Printf("Client streaming is finished ")
	if err !=nil{
		log.Fatalf("Error while recieving %v",err)
	}
	log.Printf("%v",res.Messages)
}