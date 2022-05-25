package controller

import (
	"airport/pkg/model"
	"airport/pkg/service"
)

func GetAllSeatssByDestination(origin string, destination string) []model.Seat {
	return service.GetAllSeatsByDestination(origin, destination)
}
