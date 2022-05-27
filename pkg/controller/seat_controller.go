package controller

import (
	"airport/pkg/database"
	"airport/pkg/model"
	"airport/pkg/service"
)

func GetAllSeatssByDestination(origin string, destination string) []model.Seat {
	return service.GetAllSeatsByDestination(database.GetClient(), origin, destination)
}
