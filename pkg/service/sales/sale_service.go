package sales

import (
	"airport/pkg/database"
	"airport/pkg/model"
	"airport/pkg/service/queries"
	service "airport/pkg/service/seats"
	"errors"
	"fmt"
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

	price := service.CalculateSeatPrice(service.GetSeatAvailability, seat)
	sale := model.Sale{Passenger: passenger, PassengerID: passenger.ID, SeatID: seatId, Seat: seat, Price: price, ReservationDate: time.Now()}

	//TODO: Check errors
	database.GetClient().Create(&sale)
	return sale, nil
}

func GetSale(saleID string) (model.Sale, error) {
	var sale model.Sale
	r := database.GetClient().Where("sales.id = ?", saleID).Preload(clause.Associations).Find(&sale)
	if r.Error != nil {
		return sale, errors.New("sale not found")
	}

	return sale, nil
}

func ProcessPayment(sale model.Sale, cardNumber int64, securityNumber int, expirationDate string) error {
	if !sale.Seat.IsReserved() {
		return errors.New("seat is not reserved")
	}

	if !sale.IsSaleDateNull() {
		return errors.New("sale is already paid")
	}

	// call payment api
	cardValidationFetch := queries.FakeFetch(fmt.Sprintf("api/bank/card_number=%d", cardNumber))
	cardPaymentFetch := queries.FakeFetch(fmt.Sprintf("api/payment/card_number=%d", cardNumber))

	_, err := queries.FanInFetch(cardValidationFetch, cardPaymentFetch)
	if err != nil {
		return errors.New("there was an error processing your payment")
	}

	sale.Seat.SetOccupied()
	sale.SetSaleDateAsCurrent()

	r := database.GetClient().Session(&gorm.Session{FullSaveAssociations: true}).Save(&sale)
	if r.Error != nil {
		return errors.New("error processing payment")
	}

	return nil
}
