package model

type SeatStatus string

const (
	Empty    SeatStatus = "EMPTY"
	Reserved SeatStatus = "RESERVED"
	Occupied SeatStatus = "OCCUPIED"
)

type Seat struct {
	ID           int `gorm:"primaryKey"`
	FlightID     int
	Flight       Flight `gorm:"foreignKey:ID;references:FlightID"`
	SeatLocation string
	TypeID       int
	Type         SeatType `gorm:"foreignKey:ID;references:TypeID"`
	Status       SeatStatus
}
