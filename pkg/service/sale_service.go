package service

import (
	"airport/pkg/database"
	"airport/pkg/model"
	"errors"
	"time"
)

func BookFlightSeat(seatId int) error {
	client := database.GetClient()

	var seat model.Seat
	client.Where("seats.id = ?", seatId).Find(&seat)
	if seat.ID == 0 {
		return errors.New("seat id non-existent")
	}
	if seat.Status == model.Reserved || seat.Status == model.Occupied {
		return errors.New("seat is not available")
	}
	seat.Status = model.Reserved

	//TODO: Check errors
	client.Save(seat)
	return nil
}

func SaveSale(seatID int) error {
	client := database.GetClient()

	//TODO: Add missing fields
	sale := model.Sale{SeatID: seatID, ReservationDate: time.Now(), PassengerID: 1}

	//TODO: Check errors
	client.Create(&sale)
	return nil
}
