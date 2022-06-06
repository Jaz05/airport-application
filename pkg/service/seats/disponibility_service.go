package service

func CalculateAvailability(origin string, destination string) int {
	var allSeats = GetAllSeatsByDestination(origin, destination)
	var availableSeats = GetAvailableSeatsByDestination(origin, destination)
	var allSeatsAmount = len(allSeats)
	if allSeatsAmount == 0 {
		return 0
	}
	return len(availableSeats) * 100 / allSeatsAmount
}
