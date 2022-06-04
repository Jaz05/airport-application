package service

import (
	"airport/pkg/database"
	"airport/pkg/model"
	"errors"
	"time"
)

func BookFlightSeat(seatId int) error {
	var seat model.Seat
	database.GetClient().Where("seats.id = ?", seatId).Find(&seat)
	if seat.ID == 0 {
		return errors.New("seat id non-existent")
	}
	if seat.Status == model.Reserved || seat.Status == model.Occupied {
		return errors.New("seat is not available")
	}
	seat.Status = model.Reserved

	//TODO: Check errors
	database.GetClient().Save(seat)
	return nil
}

func SaveSale(seatID int) error {
	//TODO: Add missing fields
	sale := model.Sale{SeatID: seatID, ReservationDate: time.Now(), PassengerID: 1}

	//TODO: Check errors
	database.GetClient().Create(&sale)
	return nil
}
