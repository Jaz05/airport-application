package service

import (
	"airport/pkg/model"
	"math"
)

func CalculateSeatPrice(seat model.Seat) float32 {
	var basePrice = seat.Flight.BasePrice
	var seatTypeMultiplier = seat.Type.Multiplier
	var availability = getSeatAvailability(seat.Flight.OriginID, seat.Flight.DestinationID)
	var availabilityMultiplier = calculateAvailabilityMultiplier(availability)

	// 2 decimal precision
	return float32(math.Round(float64(basePrice*seatTypeMultiplier*availabilityMultiplier)*100) / 100)
}

func calculateAvailabilityMultiplier(availability int) float32 {
	if availability >= 20 && availability < 50 {
		return 1.2
	}
	if availability < 20 {
		return 1.5
	}
	return 1

}
