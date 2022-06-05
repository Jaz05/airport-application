package sales

import (
	"airport/pkg/database"
	"airport/pkg/loader"
	"airport/pkg/model"
	"database/sql"
	"testing"
	"time"
)

func TestBookFlightSeatShouldReturnSeatWithStatusReserved(t *testing.T) {
	var flights = []model.Flight{{DestinationID: 1}}
	var seats = []model.Seat{{FlightID: 1, Status: model.Empty}}
	db := database.GetInMemoryClient()
	loader.LoadTables(db)
	db.Create(flights)
	db.Create(seats)

	seat, err := BookFlightSeat(db, 1)

	if seat.Status != model.Reserved || err != nil {
		t.Fatalf("Expected: %q, Error: %v, Got: %v", model.Reserved, err, seat.Status)
	}

}

func TestSaveSaleShouldReturnNewSale(t *testing.T) {
	var flights = []model.Flight{{DestinationID: 1}}
	var seats = []model.Seat{{FlightID: 1, Status: model.Empty}}
	db := database.GetInMemoryClient()
	loader.LoadTables(db)
	db.Create(flights)
	db.Create(seats)

	expectedPassenger := model.Passenger{ID: 1, Name: "asd", SurName: "asd", Dni: 123}
	expectedSale := model.Sale{
		ID:              1,
		PassengerID:     1,
		Passenger:       expectedPassenger,
		SeatID:          1,
		Seat:            seats[0],
		Price:           0,
		SaleDate:        sql.NullTime{},
		ReservationDate: time.Time{},
	}

	requestPassenger := model.Passenger{ID: 1, Name: "asd", SurName: "asd", Dni: 123}

	sale, err := SaveSale(db, 1, requestPassenger.Dni, requestPassenger.Name, requestPassenger.SurName)

	if expectedSale.Passenger != sale.Passenger || expectedSale.Seat != sale.Seat || err != nil {
		t.Fatalf("Expected: %v, Error: %v, Got: %v", expectedSale, err, sale)
	}

}
