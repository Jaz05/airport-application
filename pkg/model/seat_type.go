package model

import "gorm.io/gorm"

type SeatCategory int

const (
	FirstClass   SeatCategory = 1
	TouristClass SeatCategory = 2
)

type SeatType struct {
	gorm.Model
	Category   SeatCategory
	Multiplier float32
}
