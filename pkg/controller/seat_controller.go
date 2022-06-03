package controller

import (
	"airport/pkg/database"
	service "airport/pkg/service/seats"
)

type SeatResponse struct {
	Id       int
	Price    float32
	Position string
}

func GetAllSeatsByDestination(origin string, destination string) []SeatResponse {
	var seats = service.GetAllSeatsByDestination(database.GetClient(), origin, destination)
	var responseList []SeatResponse
	for _, seat := range seats {
		var price = service.CalculateSeatPrice(seat)
		var element = SeatResponse{Id: seat.ID, Position: seat.SeatLocation, Price: price}
		responseList = append(responseList, element)
	}

	return responseList
}
