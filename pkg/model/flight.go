package model

import (
	"time"
)

type Flight struct {
	ID            int `gorm:"primaryKey"`
	Capacity      int
	BasePrice     float32
	Date          time.Time
	OriginID      int
	Origin        Airport `gorm:"foreignKey:ID;references:OriginID"`
	DestinationID int
	Destination   Airport `gorm:"foreignKey:ID;references:DestinationID"`
}
