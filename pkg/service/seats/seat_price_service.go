package service

import (
	"airport/pkg/model"
	"gorm.io/gorm"
	"math"
)

func CalculateSeatPrice(client *gorm.DB, seat model.Seat) float32 {
	var basePrice = seat.Flight.BasePrice
	var seatTypeMultiplier = seat.Type.Multiplier
	var disponibility = getSeatDisponibility(client, seat.Flight.OriginID, seat.Flight.DestinationID)
	var disponibilityMultiplier = calculateDisponibilityMultiplier(disponibility)

	// 2 decimal precision
	return float32(math.Round(float64(basePrice*seatTypeMultiplier*disponibilityMultiplier)*100) / 100)
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
