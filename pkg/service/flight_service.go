package service

import (
	"airport/pkg/database"
	"airport/pkg/model"
)

func GetAllFlights() []model.Flight {
	var flights []model.Flight
	var client = database.GetClient()
	client.Find(&flights)
	return flights
}
