package server

import (
	"RecommendationService/internal/proto/pb"
	"context"
)

//func haversine(lat1, lon1, lat2, lon2 float64) float64 {
//	dLat := (lat2 - lat1) * (math.Pi / 180.0)
//	dLon := (lon2 - lon1) * (math.Pi / 180.0)
//	lat1 = lat1 * (math.Pi / 180.0)
//	lat2 = lat2 * (math.Pi / 180.0)
//
//	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
//		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)
//	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
//	return EarthRadiusKm * c
//}

type RecommendationServiceImpl struct {
	pb.RecommendationServiceServer
}

func (s *RecommendationServiceImpl) GetRecommendations(ctx context.Context, request *pb.RecommendationRequest) (*pb.RecommendationResponse, error) {
	//latitude := request.Location.Latitude
	//longitude := request.Location.Longitude
	//radiusKm := 1.0
	//apiKey := "50ce8a67-accf-4000-a8ab-57ecb6043c54"
	//
	//url := fmt.Sprintf("https://search-maps.yandex.ru/v1/nearest?apikey=%s&lat=%f&lon=%f&type=eat&lang=en_US", apiKey, latitude, longitude)
	//
	//resp, err := http.Get(url)
	//if err != nil {
	//	return nil, err
	//}
	//defer resp.Body.Close()
	//
	//var response struct {
	//	Features []struct {
	//		Properties struct {
	//			Name    string `json:"name"`
	//			Address string `json:"address"`
	//		} `json:"properties"`
	//		Geometry struct {
	//			Coordinates []float64 `json:"coordinates"`
	//		} `json:"geometry"`
	//	} `json:"features"`
	//}
	//
	//err = json.NewDecoder(resp.Body).Decode(&response)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var nearbyRestaurants []*pb.Recommendation
	//
	//for _, feature := range response.Features {
	//	distance := haversine(latitude, longitude, feature.Geometry.Coordinates[1], feature.Geometry.Coordinates[0])
	//	if distance <= radiusKm {
	//		rec := &pb.Recommendation{
	//			Name:    feature.Properties.Name,
	//			Address: feature.Properties.Address,
	//		}
	//		nearbyRestaurants = append(nearbyRestaurants, rec)
	//	}
	//}
	//
	//recommendations := pb.RecommendationResponse{
	//	Recommendations: nearbyRestaurants,
	//}
	//
	//return &recommendations, nil

	return &pb.RecommendationResponse{
		Recommendations: []*pb.Recommendation{
			&pb.Recommendation{
				Name: "test",
				Address: &pb.Location{
					Latitude:  2,
					Longitude: 3,
				},
				Distance: 15,
			},
		},
	}, nil
}
