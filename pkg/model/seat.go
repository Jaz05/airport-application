package model

import "gorm.io/gorm"

type SeatStatus string

const (
	Empty    SeatStatus = "EMPTY"
	Reserved SeatStatus = "RESERVED"
	Occupied SeatStatus = "OCCUPIED"
)

type Seat struct {
	gorm.Model
	FlightID     int
	Flight       Flight
	SeatLocation string
	TypeID       int
	Type         SeatType
	Status       SeatStatus
}
