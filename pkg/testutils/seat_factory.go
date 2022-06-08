package testutils

import "airport/pkg/model"

func NewOccupiedSeat(flightID int) model.Seat {
	var seat = model.Seat{FlightID: flightID, Status: model.Occupied}
	return seat
}

func NewEmptySeat(flightID int) model.Seat {
	var seat = model.Seat{FlightID: flightID, Status: model.Empty}
	return seat
}
