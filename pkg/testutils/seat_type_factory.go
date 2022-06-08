package testutils

import "airport/pkg/model"

// TODO: el multiplier no depende de la categoria?
func NewTouristicSeatType() model.SeatType {
	var seatType = model.SeatType{
		ID:         1,
		Category:   model.TouristClass,
		Multiplier: 1,
	}
	return seatType
}
