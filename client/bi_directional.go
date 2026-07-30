package main

import (
	"context"
	pb "grpc/protoc"
	"io"
	"log"
	"time"
)

func callSayHelloBiDirectionalStream(client pb.GreetServiceClient, names *pb.NameList) {

	log.Printf("Birectional Streaming Has Started")
	stream, err := client.SayHelloBidirectional(context.Background())
	if err != nil {
		log.Fatalf("cound not send names")
	}
	waitc := make(chan struct{})
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Some error while streaming %v", err)
			}
			log.Printf("%v", message)
		}
		close(waitc)
	}()
	for _, names := range names.Names {
		req := &pb.HelloRequest{
			Name: names,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("cannot send %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional Streaming Ended")
}
