package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Latitude  float64
	Longitude float64
}

type Recommendation struct {
	gorm.Model
	Name     string
	Location string
}
