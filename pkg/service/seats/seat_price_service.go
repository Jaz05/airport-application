package service

import "airport/pkg/model"

func CalculateSeatPrice(seat model.Seat) float32 {
	var basePrice = seat.Flight.BasePrice
	var seatTypeMultiplier = seat.Type.Multiplier
	var disponibilityMultiplier = CalculateDisponibilityMultiplier(seat)

	return basePrice * seatTypeMultiplier * disponibilityMultiplier
}

func CalculateDisponibilityMultiplier(seat model.Seat) float32 {
	return 0
}
