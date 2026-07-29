package main

import (
	"context"
	pb "grpc/protoc"
	"io"
	"log"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("streaming Started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names %v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading message %v", err)
		}
		log.Print(message)
	}
	log.Printf("streaming Finished")
}
