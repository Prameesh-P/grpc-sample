package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Prameesh-P/grpc-sample/proto"
)

func callSayHello(client pb.GrpcServiceClient) {
	ctx,cancel:= context.WithTimeout(context.Background(),time.Second)

	defer cancel()

	res,err:=client.SayHello(ctx,&pb.NoParam{})
	if err !=nil {
		log.Fatalf("could not print response.. Got an error %v",err)
	}
	log.Printf("%s",res.Message)
}