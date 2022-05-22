package model

import (
	"time"
)

type Sale struct {
	Id        int
	Passenger Passenger
	Seat      Seat
	Price     float32
	Date      time.Time
}
