package client

import (
	"LocationService/internal/proto/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
)

// client
func GetRecommendations(location *pb.Location) ([]*pb.Recommendation, error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial("127.0.0.1:50051", opts...)

	if err != nil {
		log.Fatalf("Failed to connect to recommendation service: %v", err)
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewRecommendationServiceClient(conn)
	request := &pb.RecommendationRequest{
		Location: &pb.Location{Latitude: 1456, Longitude: 212},
	}
	response, err := client.GetRecommendations(context.Background(), request)

	if err != nil {
		log.Fatalf("Failed to get recommendations: %v", err)
		grpclog.Fatalf("fail to dial: %v", err)
	}

	return response.Recommendations, nil
}
