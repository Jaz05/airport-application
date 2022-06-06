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
	Flight       Flight
	SeatLocation string
	TypeID       int
	Type         SeatType
	Status       SeatStatus
}

func (s *Seat) SetOccupied() {
	s.Status = Occupied
}

func (s *Seat) IsReserved() bool {
	return s.Status == Reserved
}
