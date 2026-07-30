package main

import (
	pb "grpc/protoc"
	"io"
	"log"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var message []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Messages: message})
		}
		if err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("Got Request with name %v", req.Name)
		message = append(message, "Hello ", req.Name)
	}
	return nil
}
