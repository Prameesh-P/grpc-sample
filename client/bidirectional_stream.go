package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Prameesh-P/grpc-sample/proto"
)

func callSayHelloBidirectionalStream(client pb.GrpcServiceClient,names *pb.NamesList){
		log.Printf("Bidirectional Streaming is started..")
		stream,err:=client.SayHelloBidirectionalStreaming(context.Background())
		if err != nil{
			log.Fatalf("could not send names..%v",err)
		}
		waitChan:=make(chan struct{})

		go func ()  {
			for {
				message,err:=stream.Recv()
				if err == io.EOF{
					break
				}
				if err !=nil{
					log.Fatalf("Error while streaming %v",err)
				}
				log.Println(message)
			}	
			close(waitChan)
		}()
		for _,name := range names.Names{
			req:=&pb.HelloRequest{
				Name: name,
			}
			if err:=stream.Send(req);err!=nil{
				log.Fatalf("could not send request %v",err)
			}
			time.Sleep(time.Second)
		}
		stream.CloseSend()
		<-waitChan
		log.Printf("Biderectional streaming finished ")

}	