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

// TODO: ISSUE: que pasa si falla el savesale? te queda el asiento reservado pero sin ninguna sale asociada
// SaveSale TODO: pasarle el body directamente? mover la logica de ver si el passenger existe a otra service que se ejecute antes?
func SaveSale(seatId int, pDni int64, pName string, pSurname string) (model.Sale, error) {
	// fetch passenger and seat
	var seat model.Seat
	var passenger model.Passenger

	c := make(chan string)
	go func() {
		database.GetClient().Where("seats.id = ?", seatId).Preload(clause.Associations).Find(&seat)
		c <- "Done"
	}()
	go func() {
		database.GetClient().Where("passengers.dni = ?", pDni).Find(&passenger)
		c <- "Done"
	}()
	for i := 0; i < 2; i++ {
		<-c
	}

	// if passenger does not exist, create new one
	if passenger.ID == 0 {
		passenger = *model.NewPassenger(pName, pSurname, pDni)
		database.GetClient().Create(&passenger)
	}

	price := service.CalculateSeatPrice(service.GetSeatAvailability, seat)
	sale := *model.NewSale(passenger.ID, passenger, seatId, seat, price)
	sale.SetReservationDateAsCurrent()

	r := database.GetClient().Create(&sale)
	if r.Error != nil {
		return sale, errors.New("error creating sale")
	}
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
