package service

import (
	"airport/pkg/model"
	"testing"
	"time"
)

func TestCalculateSeatPriceWhenDisponibilityIsLessThan20AndTypeIsEconomicAndBasePriceIs100ShouldReturn150(t *testing.T) {
	var seat = model.Seat{Flight: model.Flight{
		ID:            0,
		Capacity:      0,
		BasePrice:     100,
		Date:          time.Time{},
		OriginID:      0,
		Origin:        model.Airport{},
		DestinationID: 0,
		Destination:   model.Airport{},
	}, Type: model.SeatType{
		ID:         0,
		Category:   2,
		Multiplier: 1,
	}}

	// mock getSeatDisponibility func
	getSeatDisponibility = func(origin int, destination int) int {
		return 19
	}
	var expectedPrice float32 = 150.0

	price := CalculateSeatPrice(seat)

	if expectedPrice != price {
		t.Fatalf("Expected: %v, Got: %v", expectedPrice, price)
	}

}
