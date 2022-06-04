package service

import (
	"airport/pkg/database"
	"airport/pkg/model"
	"gorm.io/gorm/clause"
)

func GetAllFlights() []model.Flight {
	var flights []model.Flight
	database.GetClient().Preload(clause.Associations).Preload("Destination.Place").Preload("Origin.Place").Find(&flights)
	return flights
}

func GetAllFlightsByDestination(destination string) []model.Flight {
	var flights []model.Flight
	database.GetClient().Where("destination_id = ?", destination).Preload(clause.Associations).Find(&flights)
	return flights
}
