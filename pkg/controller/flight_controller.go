package controller

import (
	"airport/pkg/model"
	"airport/pkg/service"
)

func GetAllFlights() []model.Flight {
	return service.GetAllFlights()
}

func GetAllFlightsByDestination(destination string) []model.Flight {
	return service.GetAllFlightsByDestination(destination)
}
