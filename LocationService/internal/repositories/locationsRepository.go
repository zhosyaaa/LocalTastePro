package repositories

import "LocationService/internal/models"

type LocationRepo interface {
	CreateLocation(location *models.Location) (*models.Location, error)
	CreateRestaurants(rest *models.Restaurant) (*models.Restaurant, error)
}
