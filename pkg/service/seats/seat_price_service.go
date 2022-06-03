package service

import (
	"airport/pkg/model"
)

func CalculateSeatPrice(seat model.Seat) float32 {
	var basePrice = seat.Flight.BasePrice
	var seatTypeMultiplier = seat.Type.Multiplier
	var disponibility = GetSeatDisponibility(seat.Flight.OriginID, seat.Flight.DestinationID)
	var disponibilityMultiplier = calculateDisponibilityMultiplier(disponibility)

	return basePrice * seatTypeMultiplier * disponibilityMultiplier
}

func calculateDisponibilityMultiplier(disponibility int) float32 {
	if disponibility >= 20 && disponibility < 50 {
		return 1.2
	}
	if disponibility < 20 {
		return 1.5
	}
	return 1

}
