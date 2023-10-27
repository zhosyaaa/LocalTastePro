package app

import (
	"LocationService/internal/db"
	"LocationService/internal/proto/pb"
	"LocationService/internal/repositories"
	"LocationService/internal/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Run() {
	db.InitDatabase()
	repo := repositories.NewLocationRepoImpl(*db.DB)
	server := server.NewServer(*repo)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	serv := grpc.NewServer()
	pb.RegisterRecommendationServiceServer(serv, server)

	log.Println("gRPC server is running on :50051")

	if err := serv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
