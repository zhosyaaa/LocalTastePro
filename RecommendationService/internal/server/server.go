package server

import (
	"RecommendationService/internal/proto/pb"
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type RecommendationServiceImpl struct {
	pb.RecommendationServiceServer
}

func (s *RecommendationServiceImpl) GetRecommendations(ctx context.Context, request *pb.RecommendationRequest) (*pb.RecommendationResponse, error) {
	Longitude := request.Location.Longitude
	Latitude := request.Location.Latitude
	url := fmt.Sprintf("https://2gis.kz/astana/search/%D0%9F%D0%BE%D0%B5%D1%81%D1%82%D1%8C?m=%f%%2C%f%%2F16.28", Longitude, Latitude)
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

	recomendateions := make([]*pb.Recommendation, 0)
	doc.Find("._1kf6gff").Each(func(index int, rowSelection *goquery.Selection) {
		name := rowSelection.Find("._1al0wlf span").First().Text()
		stars := rowSelection.Find("._y10azs").Text()
		locate := rowSelection.Find("._1w9o2igt").Text()
		description := rowSelection.Find("._d76pv4").Text()
		rest := pb.Recommendation{Name: name, Address: locate, Stars: stars, Description: description}
		recomendateions = append(recomendateions, &rest)
	})

	return &pb.RecommendationResponse{
		Recommendations: recomendateions,
	}, nil
}
