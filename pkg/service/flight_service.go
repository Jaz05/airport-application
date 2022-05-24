package service

import (
	"airport/pkg/database"
	"airport/pkg/model"

	"gorm.io/gorm/clause"
)

func GetAllFlights() []model.Flight {
	var flights []model.Flight
	var client = database.GetClient()
	client.Preload(clause.Associations).Preload("Destination.Place").Preload("Origin.Place").Find(&flights)
	return flights
}
