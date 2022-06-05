package service

import (
	"airport/pkg/database"
	"airport/pkg/loader"
	"airport/pkg/model"
	"testing"
)

func TestBookFlightSeatShouldReturnSeatWithStatusReserved(t *testing.T) {
	var flights = []model.Flight{{DestinationID: 1}}
	var seats = []model.Seat{{FlightID: 1, Status: model.Empty}}
	db := database.GetInMemoryClient()
	loader.LoadTables(db)
	db.Create(flights)
	db.Create(seats)

	seat, err := BookFlightSeat(1)

	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	if seat.Status != model.Reserved {
		t.Fatalf("Expected: %q, Got: %v", model.Reserved, seat.Status)
	}

}
