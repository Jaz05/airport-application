package sales

import (
	"airport/pkg/model"
	"airport/pkg/testutils"
	"database/sql"
	"errors"
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

	if expectedSale.Passenger != sale.Passenger || expectedSale.Seat != sale.Seat || sale.Token != "some token" || err != nil {
		t.Fatalf("Expected: %v, Error: %v, Got: %v", expectedSale, err, sale)
	}

}

func TestGetSalesByTokenShouldReturnListOfSales(t *testing.T) {
	testutils.BeforeEach()
	var flights = []model.Flight{{DestinationID: 1}}
	testutils.MockData(flights)
	var seats = []model.Seat{{FlightID: 1, Flight: flights[0], Status: model.Empty}}
	testutils.MockData(seats)
	var passengers = []model.Passenger{{
		ID:      0,
		Name:    "mock",
		SurName: "mock",
		Dni:     1234,
	}}
	testutils.MockData(passengers)
	var sales = []model.Sale{{
		PassengerID:     1,
		Passenger:       passengers[0],
		SeatID:          1,
		Seat:            seats[0],
		Price:           0,
		SaleDate:        sql.NullTime{},
		ReservationDate: time.Time{},
		Token:           "some token",
	}}
	testutils.MockData(sales)

	var expectedSales = []model.Sale{{
		ID:              1,
		PassengerID:     1,
		Passenger:       passengers[0],
		SeatID:          1,
		Seat:            seats[0],
		Price:           0,
		SaleDate:        sql.NullTime{},
		ReservationDate: time.Time{},
		Token:           "some token",
	}}

	sales, err := GetSalesByToken("some token")

	if len(expectedSales) != 1 || expectedSales[0].Token != sales[0].Token || err != nil {
		t.Fatalf("Expected: %v, Error: %v, Got: %v", expectedSales, err, sales)
	}

}

func TestValidateSaleShouldReturnErrorIsSaleSeatIsNotReserved(t *testing.T) {
	var seat = model.Seat{
		Status: model.Empty,
	}
	var sale = model.Sale{Seat: seat}

	expectedErr := errors.New("seat is not reserved")

	err := ValidateSale(sale)
	if expectedErr.Error() != err.Error() || err == nil {
		t.Fatalf("Expected: %v, Got: %v", expectedErr, err)
	}
}

func TestValidateSaleShouldReturnErrorIsSaleSeatAlreadyHasSaleDate(t *testing.T) {
	var seat = model.Seat{
		Status: model.Reserved,
	}
	var sale = model.Sale{Seat: seat}
	sale.SetSaleDateAsCurrent()

	expectedErr := errors.New("sale is already paid")

	err := ValidateSale(sale)
	if expectedErr.Error() != err.Error() || err == nil {
		t.Fatalf("Expected: %v, Got: %v", expectedErr, err)
	}
}
