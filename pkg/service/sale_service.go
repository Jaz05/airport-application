package service

import (
	"airport/pkg/model"
	"errors"
	"gorm.io/gorm"
)

func BookFlightSeat(client *gorm.DB, seatId string) (model.Seat, error) {
	var seat model.Seat
	client.Where("seats.id = ?", seatId).Find(&seat)
	if seat.ID == 0 {
		return seat, errors.New("seat id non existent")
	}
	if seat.Status == model.Reserved || seat.Status == model.Occupied {
		return seat, errors.New("seat is already reserved")
	}
	seat.Status = model.Reserved
	client.Save(seat)
	return seat, nil
}
