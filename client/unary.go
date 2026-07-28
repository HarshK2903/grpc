package main

import (
	"context"
	pb "grpc/protoc"
	"log"
	"time"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("cannot connet to server %v", err)
	}
	log.Printf("%v", res.Message)
}
