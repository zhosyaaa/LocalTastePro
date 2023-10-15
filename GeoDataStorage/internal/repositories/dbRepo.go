package repositories

import "GeoDataStorage/internal/models"

type Repository interface {
	getRecommendations(latitude, longitude float64) ([]models.Recommendation, error)
	insertLocation(latitude, longitude float64) error
}
