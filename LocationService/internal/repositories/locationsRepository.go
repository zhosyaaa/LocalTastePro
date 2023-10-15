package repositories

import "LocationService/internal/models"

type LocationRepo interface {
	Create(location *models.Location) (*models.Location, error)
}
