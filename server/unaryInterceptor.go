package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

func loggingUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	clientAddr := "unknown"
	if p, ok := peer.FromContext(ctx); ok {
		clientAddr = p.Addr.String()
	}
	resp, err := handler(ctx, req)
	duration := time.Since(start)
	if err != nil {
		log.Printf("[unary] error method=%s client=%s duration=%s err=%v",
			info.FullMethod, clientAddr, duration, err)
	} else {
		log.Printf("[unary] done  method=%s client=%s duration=%s",
			info.FullMethod, clientAddr, duration)
	}
	return resp, err
}
