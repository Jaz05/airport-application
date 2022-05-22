package model

type SeatCategory int

const (
	FirstClass   SeatCategory = 1
	TouristClass SeatCategory = 2
)

type SeatType struct {
	Id         SeatCategory
	Multiplier float32
}
