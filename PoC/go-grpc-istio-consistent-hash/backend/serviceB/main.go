package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "./proto"
	"google.golang.org/grpc"
)

type serverB struct {
	pb.UnimplementedServiceBServer
}

func (s *serverB) HandleRequest(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	fmt.Printf("📥 Received request from pod: %s with user-id: %s\n", req.SenderPod, req.UserId)
	return &pb.Response{Message: "Received by B"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterServiceBServer(grpcServer, &serverB{})
	log.Println("🚀 Service B gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
