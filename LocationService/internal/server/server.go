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
	pb.UnimplementedLocationServiceServer
	service repositories.LocationService
}

func NewServer(service repositories.LocationService) *Server {
	return &Server{service: service}
}

func (s *Server) SendLocation(ctx context.Context, req *pb.LocationRequest) (*pb.LocationResponse, error) {
	location := &models.Location{
		Latitude:  req.Location.Latitude,
		Longitude: req.Location.Longitude,
	}
	location, err := s.service.Create(location)
	if err != nil {
		log.Fatalf("Failed to created location: %v", err)
	}
	recommendations, err := client.GetRecommendations(req.Location)
	if err != nil {
		log.Printf("Failed to get recommendations: %v", err)
	}

	response := &pb.LocationResponse{
		Id:         int32(location.ID),
		Recomendet: recommendations,
	}

	return response, nil
}
