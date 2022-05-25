package service

import (
	"airport/pkg/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllFlights(client *gorm.DB) []model.Flight {
	var flights []model.Flight
	client.Preload(clause.Associations).Preload("Destination.Place").Preload("Origin.Place").Find(&flights)
	return flights
}

func GetAllFlightsByDestination(client *gorm.DB, destination string) []model.Flight {
	var flights []model.Flight
	client.Where("destination_id = ?", destination).Preload(clause.Associations).Find(&flights)
	return flights
}
