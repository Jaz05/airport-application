package model

import (
	"time"

	"gorm.io/gorm"
)

type Flight struct {
	gorm.Model
	Capacity      int
	BasePrice     float32
	Date          time.Time
	OriginID      int
	Origin        Airport
	DestinationID int
	Destination   Airport
}
