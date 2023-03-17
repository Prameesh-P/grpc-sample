package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Prameesh-P/grpc-sample/proto"
)

func callSayHelloServerStream(client pb.GrpcServiceClient,names *pb.NamesList){
	log.Printf("Streaming started..")
	stream,err:=client.SayHelloServerSideStreaming(context.Background(),names )
	if err !=nil {
		log.Fatalf("Could not send names got an error %v",err)
	}
	for {
		message,err:= stream.Recv()
		if err == io.EOF{
			break
		}
		if err!=nil {
			log.Fatalf("Error got while streaming %v",err)
		}
			log.Println(message)
	}
	log.Println("Streaming finished..!!")
}