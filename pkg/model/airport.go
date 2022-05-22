package model

type AirportStatus string

const (
	Open   AirportStatus = "OPEN"
	Closed AirportStatus = "CLOSED"
)

type Airport struct {
	Id     int
	Name   string
	Place  Place
	Status AirportStatus
}
