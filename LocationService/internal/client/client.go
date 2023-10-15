package client

import (
	"LocationService/internal/proto/pb"
	"context"
	"google.golang.org/grpc"
	"log"
)

func SendLocationToRecommendationService(location *pb.Location) ([]*pb.Recommendation, error) {
	recommendationConn, err := grpc.Dial("recommendationservice:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to recommendation service: %v", err)
		return nil, err
	}
	defer recommendationConn.Close()

	recommendationClient := pb.NewRecommendationServiceClient(recommendationConn)

	ctx := context.Background()

	recommendations, err := recommendationClient.GetRecommendations(ctx, &pb.RecommendationRequest{
		Location: location,
	})
	if err != nil {
		log.Fatalf("Failed to get recommendations: %v", err)
		return nil, err
	}

	return recommendations.Recommendations, nil
}

func GetRecommendations(location *pb.Location) ([]*pb.Recommendation, error) {
	recommendations, err := SendLocationToRecommendationService(location)
	if err != nil {
		return nil, err
	}

	return recommendations, nil
}
