package service

import (
	"gorm.io/gorm"
)

func CalculateDisponibility(client *gorm.DB, origin string, destination string) int {
	var allSeats = GetAllSeatsByDestination(client, origin, destination)
	var availalbleSeats = GetAvailableSeatsByDestination(client, origin, destination)
	var allSeatsAmount = len(allSeats)
	if allSeatsAmount == 0 {
		return 0
	}
	return len(availalbleSeats) * 100 / allSeatsAmount
}
