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
	if seat.Status == model.Occupied || seat.Status == model.Reserved {
		return errors.New("seat is not available")
	}

	seat.Status = model.Reserved

	//TODO: Check errors
	// TODO: use update instead of save to prevent possible creation of a new seat
	database.GetClient().Save(seat)
	return nil
}

func SaveSale(seatId int, pDni int64, pName string, pSurname string) (model.Sale, error) {
	// fetch passenger and seat
	var seat model.Seat
	var passenger model.Passenger

	// TODO: usar go routines y canales
	database.GetClient().Where("seats.id = ?", seatId).Find(&seat)
	database.GetClient().Where("passengers.dni = ?", pDni).Find(&passenger)

	// if passenger doesnt not exist, create new one
	if passenger.ID == 0 {
		passenger.Name = pName
		passenger.SurName = pSurname
		passenger.Dni = pDni
		database.GetClient().Create(&passenger)
	}

	//TODO: calcular precio
	sale := model.Sale{Passenger: passenger, PassengerID: passenger.ID, SeatID: seatId, Seat: seat, Price: 0, ReservationDate: time.Now()}

	//TODO: Check errors
	database.GetClient().Create(&sale)
	return sale, nil
}
