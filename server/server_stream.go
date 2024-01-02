package main

import (
	"log"
	"time"

	pb "github.com/SonLPH/basic-go-grpc/proto"
)

func (s *helloServer) SayHelloServerStream(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamServer) error {
	log.Printf("Got request with names: %v", req.Name)
	for _, name := range req.Name {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
