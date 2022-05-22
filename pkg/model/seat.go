package model

type SeatStatus string

const (
	Empty    SeatStatus = "EMPTY"
	Reserved SeatStatus = "RESERVED"
	Occupied SeatStatus = "OCCUPIED"
)

type Seat struct {
	Id           int
	Flight       Flight
	SeatLocation string
	Type         SeatType
	Status       SeatStatus
}
