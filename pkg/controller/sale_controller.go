package controller

import (
	"airport/pkg/database"
	"airport/pkg/model"
	"airport/pkg/service"
)

func BookFlightSeat(seatId string) (model.Seat, error) {
	return service.BookFlightSeat(database.GetClient(), seatId)
}
