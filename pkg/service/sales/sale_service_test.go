package sales

import (
	"airport/pkg/model"
	"airport/pkg/testutils"
	"database/sql"
	"testing"
	"time"
)

func TestBookFlightSeatShouldReturnSeatWithStatusReserved(t *testing.T) {
	testutils.BeforeEach()
	testutils.MockData([]model.Flight{{DestinationID: 1}})
	testutils.MockData([]model.Seat{{FlightID: 1, Status: model.Empty}})

	seat, err := BookFlightSeat(1)

	if seat.Status != model.Reserved || err != nil {
		t.Fatalf("Expected: %q, Error: %v, Got: %v", model.Reserved, err, seat.Status)
	}

}

func TestSaveSaleShouldReturnNewSale(t *testing.T) {
	testutils.BeforeEach()
	var flights = []model.Flight{{DestinationID: 1}}
	testutils.MockData(flights)
	var seats = []model.Seat{{FlightID: 1, Flight: flights[0], Status: model.Empty}}
	testutils.MockData(seats)

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

	sale, err := SaveSale(1, requestPassenger.Dni, requestPassenger.Name, requestPassenger.SurName, "some token")

	if expectedSale.Passenger != sale.Passenger || expectedSale.Seat != sale.Seat || err != nil {
		t.Fatalf("Expected: %v, Error: %v, Got: %v", expectedSale, err, sale)
	}

}
