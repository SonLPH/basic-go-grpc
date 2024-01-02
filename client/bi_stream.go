package main

import (
	"context"
	"io"
	"log"

	pb "github.com/SonLPH/basic-go-grpc/proto"
)

func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional streaming started")
	stream, err := client.SayHelloBidirectionalStream(context.Background())

	if err != nil {
		log.Fatalf("Failed to call SayHelloBidirectionalStream: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Failed to receive a message: %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Name {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Printf("Error while sending request: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Birectional streaming finished!")
}
