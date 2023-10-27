package main

import (
	"RecommendationService/internal/proto/pb"
	"RecommendationService/internal/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	serv := grpc.NewServer()
	pb.RegisterRecommendationServiceServer(serv, &server.RecommendationServiceImpl{})

	log.Println("gRPC server is running on :50051")

	if err := serv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
