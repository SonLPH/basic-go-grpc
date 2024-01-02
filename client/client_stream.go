package main

import (
	"context"
	"log"
	"time"

	pb "github.com/SonLPH/basic-go-grpc/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client streaming started")
	stream, err := client.SayHelloClientStream(context.Background())
	if err != nil {
		log.Fatalf("Failed to call SayHelloClientStream: %v", err)
	}

	for _, name := range names.Name {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Failed to send a request: %v", err)
		}
		log.Printf("Request sent with name: %v", name)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client streaming finished with response: %v", res.Message)
	if err != nil {
		log.Fatalf("Failed to receive a response: %v", err)
	}
}
