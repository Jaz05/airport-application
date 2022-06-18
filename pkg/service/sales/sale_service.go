package sales

import (
	"airport/pkg/database"
	"airport/pkg/model"
	service "airport/pkg/service/seats"
	"errors"

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

	r := database.GetClient().Model(&seat).Update("status", model.Reserved)
	if r.Error != nil {
		return seat, errors.New("error updating seat status")
	}
	return seat, nil
}

// TODO: ISSUE: que pasa si falla el savesale? te queda el asiento reservado pero sin ninguna sale asociada

func SaveSale(seatId int, pDni int64, pName string, pSurname string, token string) (model.Sale, error) {
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
	sale := model.NewSale(passenger.ID, passenger, seatId, seat, price)
	sale.SetReservationDateAsCurrent()
	sale.Token = token

	r := database.GetClient().Create(sale)
	if r.Error != nil {
		return *sale, errors.New("error creating sale")
	}
	return *sale, nil
}

func GetSalesByToken(token string) ([]model.Sale, error) {
	var sales []model.Sale
	r := database.GetClient().Where("sales.token = ?", token).Preload(clause.Associations).Find(&sales)
	if r.Error != nil {
		return sales, errors.New("sale not found")
	}

	return sales, nil
}

func ValidateSale(sale model.Sale) error {
	if !sale.Seat.IsReserved() {
		return errors.New("seat is not reserved")
	}

	if !sale.IsSaleDateNull() {
		return errors.New("sale is already paid")
	}

	return nil

}

func FulfillSale(sale model.Sale) error {
	sale.Seat.SetOccupied()
	sale.SetSaleDateAsCurrent()

	r := database.GetClient().Session(&gorm.Session{FullSaveAssociations: true}).Save(&sale)
	if r.Error != nil {
		return errors.New("error processing payment")
	}
	return nil
}
