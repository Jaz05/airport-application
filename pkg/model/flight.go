package model

import (
	"time"
)

type Flight struct {
	Id        int
	Capacity  int
	BasePrice float32
	Date      time.Time
	Origin    Place
	Destination Place
}
