package model

import (
	"time"
)

type Sale struct {
	ID              int `gorm:"primaryKey"`
	PassengerID     int
	Passenger       Passenger `gorm:"foreignKey:ID;references:PassengerID"`
	SeatID          int
	Seat            Seat `gorm:"foreignKey:ID;references:SeatID"`
	Price           float32
	SaleDate        time.Time
	ReservationDate time.Time
}
