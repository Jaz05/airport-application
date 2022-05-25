package controller

import (
	"airport/pkg/database"
	"airport/pkg/model"
	"airport/pkg/service"
)

func GetAllFlights() []model.Flight {
	return service.GetAllFlights(database.GetClient())
}

func GetAllFlightsByDestination(destination string) []model.Flight {
	return service.GetAllFlightsByDestination(database.GetClient(), destination)
}
