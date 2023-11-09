package server

import (
	"RecommendationService/internal/models"
	"RecommendationService/internal/proto/pb"
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
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
	//https://www.google.com/maps/search/%D0%A0%D0%B5%D1%81%D1%82%D0%BE%D1%80%D0%B0%D0%BD%D1%8B/@51.1024504,71.4091434,14z/data=!3m1!4b1!4m7!2m6!3m5!1z0KDQtdGB0YLQvtGA0LDQvdGL!2s51.103403,+71.428386!4m2!1d71.428386!2d51.1034028?hl=ru&entry=ttu
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

func Parse() ([]models.Restaurant, error) {
	url := "https://2gis.kz/astana/search/%D0%9F%D0%BE%D0%B5%D1%81%D1%82%D1%8C?m=71.420362%2C51.089544%2F16.28"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("HTTP request returned status: %d", res.StatusCode)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("._1kf6gff").Each(func(index int, rowSelection *goquery.Selection) {
		name := rowSelection.Find("._1al0wlf span").First().Text()
		stars := rowSelection.Find("._y10azs").Text()
		locate := rowSelection.Find("._1w9o2igt").Text()
		description := rowSelection.Find("._d76pv4").Text()
		fmt.Println(index, " - ", "name: ", name, ", start: ", stars, ", locate: ", locate, ", description: ", description, "\n")
	})
	return nil, nil
}
