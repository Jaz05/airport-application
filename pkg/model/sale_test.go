package model

import (
	"database/sql"
	"testing"
	"time"
)

func TestSaleConstructorShouldReturnNewSale(t *testing.T) {
	expectedSale := Sale{
		PassengerID:     1,
		Passenger:       Passenger{},
		SeatID:          1,
		Seat:            Seat{},
		Price:           0,
		SaleDate:        sql.NullTime{},
		ReservationDate: time.Time{},
	}

	sale := *NewSale(1, Passenger{}, 1, Seat{}, 0)

	if expectedSale != sale {
		t.Fatalf("Expected: %v, Got: %v", expectedSale, sale)
	}

}
