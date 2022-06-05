package sales

import (
	"airport/pkg/model"
	"errors"
	"gorm.io/gorm"
	"time"
)

func BookFlightSeat(client *gorm.DB, seatId int) (model.Seat, error) {
	var seat model.Seat
	client.Where("seats.id = ?", seatId).Find(&seat)
	if seat.ID == 0 {
		return seat, errors.New("seat id non-existent")
	}
	if seat.Status == model.Occupied || seat.Status == model.Reserved {
		return seat, errors.New("seat is not available")
	}

	seat.Status = model.Reserved

	//TODO: Check errors
	// TODO: use update instead of save to prevent possible creation of a new seat
	client.Save(seat)
	return seat, nil
}

func SaveSale(client *gorm.DB, seatId int, pDni int64, pName string, pSurname string) (model.Sale, error) {
	// fetch passenger and seat
	var seat model.Seat
	var passenger model.Passenger

	// TODO: usar go routines y canales
	client.Where("seats.id = ?", seatId).Find(&seat)
	client.Where("passengers.dni = ?", pDni).Find(&passenger)

	// if passenger doesnt not exist, create new one
	if passenger.ID == 0 {
		passenger.Name = pName
		passenger.SurName = pSurname
		passenger.Dni = pDni
		client.Create(&passenger)
	}

	//TODO: calcular precio
	sale := model.Sale{Passenger: passenger, PassengerID: passenger.ID, SeatID: seatId, Seat: seat, Price: 0, ReservationDate: time.Now()}

	//TODO: Check errors
	client.Create(&sale)
	return sale, nil
}
