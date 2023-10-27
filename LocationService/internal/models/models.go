package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Latitude  float64
	Longitude float64
}

type Restaurant struct {
	gorm.Model
	Name           string
	Address        Location
	Distance       float64
	UserLocation   Location `gorm:"foreignKey:UserLocationID"`
	UserLocationID uint
}
