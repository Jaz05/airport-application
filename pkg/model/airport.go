package model

import (
	"gorm.io/gorm"
)

type AirportStatus string

const (
	Open   AirportStatus = "OPEN"
	Closed AirportStatus = "CLOSED"
)

type Airport struct {
	gorm.Model
	Name    string
	PlaceID int
	Place   Place
	Status  AirportStatus
}
