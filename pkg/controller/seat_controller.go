package controller

import (
	"airport/pkg/database"
	"airport/pkg/model"
	service "airport/pkg/service/seats"
)

/*
func CalculateSeatPrice(origin string, destination string) []model.Seat {
	return service.GetAllSeatsByDestination(database.GetClient(), origin, destination)
}
*/

func GetAllSeatsByDestination(origin string, destination string) []model.Seat {
	return service.GetAllSeatsByDestination(database.GetClient(), origin, destination)
}
