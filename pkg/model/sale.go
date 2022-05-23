package model

import (
	"time"

	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	PassengerID     int
	Passenger       Passenger
	SeatID          int
	Seat            Seat
	Price           float32
	SaleDate        time.Time
	ReservationDate time.Time
}
