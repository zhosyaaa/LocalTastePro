package repositories

import (
	"LocationService/internal/models"
	"gorm.io/gorm"
)

type LocationRepoImpl struct {
	db gorm.DB
}

func NewLocationRepoImpl(db gorm.DB) *LocationRepoImpl {
	return &LocationRepoImpl{db: db}
}
func (l LocationRepoImpl) CreateLocation(location *models.Location) (*models.Location, error) {
	if err := l.db.Create(location).Error; err != nil {
		return nil, err
	}
	return location, nil
}

func (l LocationRepoImpl) CreateRestaurants(rest *models.Restaurant) (*models.Restaurant, error) {
	if err := l.db.Create(rest).Error; err != nil {
		return nil, err
	}
	return rest, nil
}
