package main

import (
	"context"
	"io"
	"log"

	pb "github.com/SonLPH/basic-go-grpc/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Calling SayHelloServerStream RPC...")
	stream, err := client.SayHelloServerStream(context.Background(), names)
	if err != nil {
		log.Fatalf("Failed to call SayHelloServerStream: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive a message: %v", err)
		}
		log.Printf("Response from server: %s", message.Message)
	}
	log.Printf("Finished!")
}
