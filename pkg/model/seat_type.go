package model

type SeatCategory int

const (
	FirstClass   SeatCategory = 1
	TouristClass SeatCategory = 2
)

type SeatType struct {
	ID         int `gorm:"primaryKey"`
	Category   SeatCategory
	Multiplier float32
}
