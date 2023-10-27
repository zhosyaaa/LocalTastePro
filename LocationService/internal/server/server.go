package server

import (
	"LocationService/internal/client"
	"LocationService/internal/models"
	"LocationService/internal/proto/pb"
	"LocationService/internal/repositories"
	"context"
	"log"
)

type Server struct {
	pb.RecommendationServiceServer
	service repositories.LocationRepoImpl
}

func NewServer(service repositories.LocationRepoImpl) *Server {
	return &Server{service: service}
}

func (s *Server) SendLocation(ctx context.Context, req *pb.LocationRequest) (*pb.LocationResponse, error) {
	location := &models.Location{
		Latitude:  req.Location.Latitude,
		Longitude: req.Location.Longitude,
	}
	location, err := s.service.CreateLocation(location)
	if err != nil {
		log.Fatalf("Failed to created location: %v", err)
		return nil, err
	}
	recommendations, err := client.GetRecommendations(req.Location)
	if err != nil {
		log.Printf("Failed to get recommendations: %v", err)
		return nil, err
	}

	for i := 0; i < len(recommendations); i++ {
		restaurant := &models.Restaurant{
			Name: recommendations[i].Name,
			Address: models.Location{
				Latitude:  recommendations[i].Address.Latitude,
				Longitude: recommendations[i].Address.Longitude,
			},
			Distance: float64(recommendations[i].Distance),
		}
		s.service.CreateRestaurants(restaurant)
	}

	response := &pb.LocationResponse{
		Location:   req.Location,
		Recomendet: recommendations,
	}

	return response, nil
}
