package main

import (
	pb "grpc/protoc"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start server %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(loggingUnaryInterceptor),
	)
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server started at %v", list.Addr())

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to start grpc server %v", err)

	}
}
