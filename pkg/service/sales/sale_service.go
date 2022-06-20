package sales

import (
	"airport/pkg/database"
	"airport/pkg/dto"
	"airport/pkg/model"
	service "airport/pkg/service/seats"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func bookFlightSeat(seatId int) (model.Seat, error) {
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

func saveSale(seatId int, pDni int64, pName string, pSurname string, token string) (model.Sale, error) {
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

// TODO: deberia recibir un objeto de negocio y no un DTO

func CreateSales(salesBody dto.SalesRequestBody, token string) ([]model.Sale, error) {

	/*reps: how many sales are we doing*/
	reps := len(salesBody.Sales)
	/*reps: each sales returns a value through this seatsChannel*/
	seatsChannel := make(chan model.Seat)
	/*reps: if an error happens it returns an error, to this errorsChannel, instead */
	errorsChannel := make(chan error)

	// TODO: si un solo asiento se reserva queda reservado y el otro no (implementar rollback)
	// TODO: validar que no tengan el mismo seat id
	for i := 0; i < reps; i++ {
		body := salesBody.Sales[i]
		go func() {
			seat, err := bookFlightSeat(body.SeatId)
			if err != nil {
				errorsChannel <- err
			}
			seatsChannel <- seat
		}()
	}

	/*we expect an amount of answers equal to the length of sales*/
	/*we throw an error if at least one of them fails*/
	for i := 0; i < reps; i++ {
		select {
		case <-seatsChannel:
		case err := <-errorsChannel:
			return nil, err
		}
	}

	salesChannel := make(chan model.Sale)
	for i := 0; i < reps; i++ {
		body := salesBody.Sales[i]
		go func() {
			sale, err := saveSale(body.SeatId, body.Dni, body.Name, body.Surname, token)
			if err != nil {
				errorsChannel <- err
			}
			salesChannel <- sale
		}()
	}

	var sales []model.Sale
	for i := 0; i < reps; i++ {
		select {
		case sale := <-salesChannel:
			sales = append(sales, sale)
		case err := <-errorsChannel:
			return nil, err
		}
	}

	return sales, nil
}
