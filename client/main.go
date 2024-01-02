package main

import (
	"log"

	pb "github.com/SonLPH/basic-go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Name: []string{"Son", "Hoang", "Le"},
	}

	// Unary RPC
	// callSayHello(client)
	// Server Streaming RPC
	// callSayHelloServerStream(client, names)
	// Client Streaming RPC
	// callSayHelloClientStream(client, names)
	// Bidirectional Streaming RPC
	callSayHelloBidirectionalStream(client, names)
}
