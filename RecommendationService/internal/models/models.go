package models

type Location struct {
	Latitude  float64
	Longitude float64
}

type Restaurant struct {
	Name     string
	Address  Location
	Distance float64
}
