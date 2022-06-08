package service

import (
	"airport/pkg/model"
	"airport/pkg/testutils"
	"testing"
	"time"
)

var seat = model.Seat{Flight: model.Flight{
	ID:            0,
	Capacity:      0,
	BasePrice:     100,
	Date:          time.Time{},
	OriginID:      1,
	Origin:        model.Airport{},
	DestinationID: 2,
	Destination:   model.Airport{},
}, Type: model.SeatType{
	ID:         0,
	Category:   2,
	Multiplier: 1,
}}

func TestWithMockCalculateSeatPriceWhenAvailabilityIsLessThan20AndTypeIsEconomicAndBasePriceIs100ShouldReturn150(t *testing.T) {
	testutils.BeforeEach()

	var expectedPrice float32 = 150.0

	price := CalculateSeatPrice(testutils.MockGetSeatAvailability(19), seat)

	if expectedPrice != price {
		t.Fatalf("Expected: %v, Got: %v", expectedPrice, price)
	}

}

func TestCalculateSeatPriceWhenAvailabilityIsMoreThan20ButLessThan50AndTypeIsEconomicAndBasePriceIs100ShouldReturn120(t *testing.T) {
	testutils.BeforeEach()

	var expectedPrice float32 = 120.0

	price := CalculateSeatPrice(testutils.MockGetSeatAvailability(21), seat)

	if expectedPrice != price {
		t.Fatalf("Expected: %v, Got: %v", expectedPrice, price)
	}

}

func TestCalculateSeatPriceWhenAvailabilityIsMoreThan50AndTypeIsEconomicAndBasePriceIs100ShouldReturn100(t *testing.T) {
	testutils.BeforeEach()

	// mock getSeatAvailability func

	var expectedPrice float32 = 100.0

	price := CalculateSeatPrice(testutils.MockGetSeatAvailability(51), seat)

	if expectedPrice != price {
		t.Fatalf("Expected: %v, Got: %v", expectedPrice, price)
	}

}

func TestCalculateSeatPriceWhenAvailabilityIs16PercentAndTypeIsEconomicAndBasePriceIs100ShouldReturn150(t *testing.T) {
	var flights = []model.Flight{{OriginID: 1, DestinationID: 2}}
	var types = []model.SeatType{testutils.NewTouristicSeatType()}

	// 5 occupied, 1 empty equals 16.6% availability
	var seats = []model.Seat{
		testutils.NewOccupiedSeat(1),
		testutils.NewOccupiedSeat(1),
		testutils.NewOccupiedSeat(1),
		testutils.NewOccupiedSeat(1),
		testutils.NewOccupiedSeat(1),
		testutils.NewEmptySeat(1),
	}

	testutils.BeforeEach()
	testutils.MockData(flights)
	testutils.MockData(types)
	testutils.MockData(seats)

	var expectedPrice float32 = 150.0

	price := CalculateSeatPrice(GetSeatAvailability, seat)

	if expectedPrice != price {
		t.Fatalf("Expected: %v, Got: %v", expectedPrice, price)
	}

}

func TestCalculateSeatPriceWhenAvailabilityIs20PercentAndTypeIsEconomicAndBasePriceIs100ShouldReturn120(t *testing.T) {
	var flights = []model.Flight{{OriginID: 1, DestinationID: 2}}
	var types = []model.SeatType{testutils.NewTouristicSeatType()}

	// 4 occupied, 1 empty equals 20% availability
	var seats = []model.Seat{
		testutils.NewOccupiedSeat(1),
		testutils.NewOccupiedSeat(1),
		testutils.NewOccupiedSeat(1),
		testutils.NewOccupiedSeat(1),
		testutils.NewEmptySeat(1),
	}

	testutils.BeforeEach()
	priceMap = make(map[Route]int)
	testutils.MockData(flights)
	testutils.MockData(types)
	testutils.MockData(seats)

	var expectedPrice float32 = 120.0

	price := CalculateSeatPrice(GetSeatAvailability, seat)

	if expectedPrice != price {
		t.Fatalf("Expected: %v, Got: %v", expectedPrice, price)
	}

}

func TestCalculateSeatPriceWhenAvailabilityIs25PercentAndTypeIsEconomicAndBasePriceIs100ShouldReturn120(t *testing.T) {
	var flights = []model.Flight{{OriginID: 1, DestinationID: 2}}
	var types = []model.SeatType{testutils.NewTouristicSeatType()}

	// 3 occupied, 1 empty equals 25% availability
	var seats = []model.Seat{
		testutils.NewOccupiedSeat(1),
		testutils.NewOccupiedSeat(1),
		testutils.NewOccupiedSeat(1),
		testutils.NewEmptySeat(1),
	}

	testutils.BeforeEach()
	priceMap = make(map[Route]int)
	testutils.MockData(flights)
	testutils.MockData(types)
	testutils.MockData(seats)

	var expectedPrice float32 = 120.0

	price := CalculateSeatPrice(GetSeatAvailability, seat)

	if expectedPrice != price {
		t.Fatalf("Expected: %v, Got: %v", expectedPrice, price)
	}

}

func TestCalculateSeatPriceWhenAvailabilityIs50PercentAndTypeIsEconomicAndBasePriceIs100ShouldReturn100(t *testing.T) {
	var flights = []model.Flight{{OriginID: 1, DestinationID: 2}}
	var types = []model.SeatType{testutils.NewTouristicSeatType()}

	// 3 occupied, 1 empty equals 25% availability
	var seats = []model.Seat{
		testutils.NewOccupiedSeat(1),
		testutils.NewOccupiedSeat(1),
		testutils.NewOccupiedSeat(1),
		testutils.NewEmptySeat(1),
		testutils.NewEmptySeat(1),
		testutils.NewEmptySeat(1),
	}

	testutils.BeforeEach()
	priceMap = make(map[Route]int)
	testutils.MockData(flights)
	testutils.MockData(types)
	testutils.MockData(seats)

	var expectedPrice float32 = 100.0

	price := CalculateSeatPrice(GetSeatAvailability, seat)

	if expectedPrice != price {
		t.Fatalf("Expected: %v, Got: %v", expectedPrice, price)
	}

}
