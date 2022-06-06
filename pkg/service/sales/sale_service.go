package sales

import (
	"airport/pkg/database"
	"airport/pkg/model"
	service "airport/pkg/service/seats"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func BookFlightSeat(seatId int) (model.Seat, error) {
	var seat model.Seat
	database.GetClient().Where("seats.id = ?", seatId).Find(&seat)

	if seat.ID == 0 {
		return seat, errors.New("seat id non-existent")
	}
	if seat.Status == model.Occupied || seat.Status == model.Reserved {
		return seat, errors.New("seat is not available")
	}

	//TODO: Check errors
	database.GetClient().Model(&seat).Update("status", model.Reserved)
	return seat, nil
}

// SaveSale TODO: pasarle el body directamente? mover la logica de ver si el passenger existe a otra service que se ejecute antes?
func SaveSale(seatId int, pDni int64, pName string, pSurname string) (model.Sale, error) {
	// fetch passenger and seat
	var seat model.Seat
	var passenger model.Passenger

	// TODO: usar go routines y canales
	database.GetClient().Where("seats.id = ?", seatId).Preload(clause.Associations).Find(&seat)
	database.GetClient().Where("passengers.dni = ?", pDni).Find(&passenger)

	// if passenger doesnt not exist, create new one
	if passenger.ID == 0 {
		passenger.Name = pName
		passenger.SurName = pSurname
		passenger.Dni = pDni
		database.GetClient().Create(&passenger)
	}

	price := service.CalculateSeatPrice(seat)
	sale := model.Sale{Passenger: passenger, PassengerID: passenger.ID, SeatID: seatId, Seat: seat, Price: price, ReservationDate: time.Now()}

	//TODO: Check errors
	database.GetClient().Create(&sale)
	return sale, nil
}

func GetSale(saleID string) model.Sale {
	var sale model.Sale
	database.GetClient().Where("sales.id = ?", saleID).Preload(clause.Associations).Find(&sale)

	return sale
}

func ProcessPayment(sale model.Sale, CardNumber int64, SecurityNumber int, ExpirationDate string) {
	sale.Seat.SetOccupied()
	database.GetClient().Session(&gorm.Session{FullSaveAssociations: true}).Save(&sale)
}
