package main

import (
	pb "grpc/protoc"
	"io"
	"log"
	"time"
)

func (s *helloServer) SayHelloBidirectional(stream pb.GreetService_SayHelloBidirectionalServer) error {
	log.Printf("Server Streaming Started")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Connot Get From client %v", err)
		}
		log.Printf("Got the Request From %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
