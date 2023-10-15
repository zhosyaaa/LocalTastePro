package main

import (
	"LocationService/internal/db"
	"LocationService/internal/proto/pb"
	"LocationService/internal/repositories"
	"LocationService/internal/server"
	"google.golang.org/grpc"
	_ "google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()

	locationService := repositories.NewLocationService(*db.DB)

	locationServer := server.NewServer(*locationService)
	pb.RegisterLocationServiceServer(grpcServer, locationServer)

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("gRPC server listening on port 50051")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
