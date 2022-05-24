package model

import (
	"time"
)

type Sale struct {
	ID              int `gorm:"primaryKey"`
	PassengerID     int
	Passenger       Passenger
	SeatID          int
	Seat            Seat
	Price           float32
	SaleDate        time.Time
	ReservationDate time.Time
}
