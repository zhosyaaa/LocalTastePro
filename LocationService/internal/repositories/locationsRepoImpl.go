package repositories

import (
	"LocationService/internal/models"
	"gorm.io/gorm"
)

type LocationService struct {
	db gorm.DB
}

func NewLocationService(db gorm.DB) *LocationService {
	return &LocationService{db: db}
}

func (l LocationService) Create(location *models.Location) (*models.Location, error) {
	if err := l.db.Create(location).Error; err != nil {
		return nil, err
	}
	return location, nil
}
