package main

import (
	pb "grpc/protoc"
	"log"
	"time"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Got The list %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			log.Fatalf("Cannot Send Message %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
