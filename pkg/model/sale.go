package model

import (
	"database/sql"
	"time"
)

type Sale struct {
	ID              int `gorm:"primaryKey"`
	PassengerID     int
	Passenger       Passenger
	SeatID          int
	Seat            Seat
	Price           float32
	SaleDate        sql.NullTime
	ReservationDate time.Time
}

func (s *Sale) SetSaleDateAsCurrent() {
	s.SaleDate = sql.NullTime{Time: time.Now(), Valid: true}
}

func (s *Sale) IsSaleDateNull() bool {
	return !s.SaleDate.Valid
}
