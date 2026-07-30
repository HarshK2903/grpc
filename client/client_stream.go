package main

import (
	"context"
	pb "grpc/protoc"
	"log"
	"time"
)

func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NameList) {

	log.Printf("Client Stearming Started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Couldn't Send Names")
	}
	for _, name1 := range names.Names {
		req := &pb.HelloRequest{
			Name: name1,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending %v", err)
		}
		log.Printf("Sent the request with name %v", name1)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("Client Streaming Finished")
	if err != nil {
		log.Fatalf("error while reciving %v", err)
	}
	log.Printf("recive Meassage :- %v", res.Messages)
}
