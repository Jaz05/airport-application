package service

import (
	"airport/pkg/database"
	"airport/pkg/loader"
	"airport/pkg/model"
	"gorm.io/gorm"
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

func TestWithMockCalculateSeatPriceWhenDisponibilityIsLessThan20AndTypeIsEconomicAndBasePriceIs100ShouldReturn150(t *testing.T) {
	db := database.GetInMemoryClient()
	// mock getSeatDisponibility func
	getSeatDisponibility = func(client *gorm.DB, origin int, destination int) int {
		return 19
	}
	var expectedPrice float32 = 150.0

	price := CalculateSeatPrice(db, seat)

	if expectedPrice != price {
		t.Fatalf("Expected: %v, Got: %v", expectedPrice, price)
	}

}

func TestCalculateSeatPriceWhenDisponibilityIs20PercentAndTypeIsEconomicAndBasePriceIs100ShouldReturn150(t *testing.T) {
	var flights = []model.Flight{{OriginID: 1, DestinationID: 2}}
	var types = []model.SeatType{{
		ID:         1,
		Category:   2,
		Multiplier: 1,
	}}

	// 4 occupied, 1 empty equals 20% disponibility
	var seats = []model.Seat{{FlightID: 1, Status: "OCCUPIED"},
		{FlightID: 1, Status: "OCCUPIED"},
		{FlightID: 1, Status: "OCCUPIED"},
		{FlightID: 1, Status: "OCCUPIED"},
		{FlightID: 1, Status: "OCCUPIED"},
		{FlightID: 1, Status: "EMPTY"}}

	db := database.GetInMemoryClient()
	loader.LoadTables(db)
	db.Create(flights)
	db.Create(types)
	db.Create(seats)

	var expectedPrice float32 = 150.0

	price := CalculateSeatPrice(db, seat)

	if expectedPrice != price {
		t.Fatalf("Expected: %v, Got: %v", expectedPrice, price)
	}

}

func TestCalculateSeatPriceWhenDisponibilityIsMoreThan20ButLessThan50AndTypeIsEconomicAndBasePriceIs100ShouldReturn120(t *testing.T) {
	db := database.GetInMemoryClient()
	// mock getSeatDisponibility func
	getSeatDisponibility = func(client *gorm.DB, origin int, destination int) int {
		return 21
	}
	var expectedPrice float32 = 120.0

	price := CalculateSeatPrice(db, seat)

	if expectedPrice != price {
		t.Fatalf("Expected: %v, Got: %v", expectedPrice, price)
	}

}

func TestCalculateSeatPriceWhenDisponibilityIsMoreThan50AndTypeIsEconomicAndBasePriceIs100ShouldReturn100(t *testing.T) {
	db := database.GetInMemoryClient()
	// mock getSeatDisponibility func
	getSeatDisponibility = func(client *gorm.DB, origin int, destination int) int {
		return 51
	}
	var expectedPrice float32 = 100.0

	price := CalculateSeatPrice(db, seat)

	if expectedPrice != price {
		t.Fatalf("Expected: %v, Got: %v", expectedPrice, price)
	}

}
