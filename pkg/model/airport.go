package model

type AirportStatus string

const (
	Open   AirportStatus = "OPEN"
	Closed AirportStatus = "CLOSED"
)

type Airport struct {
	ID      int `gorm:"primaryKey"`
	Name    string
	PlaceID int
	Place   Place `gorm:"foreignKey:ID;references:PlaceID"`
	Status  AirportStatus
}
