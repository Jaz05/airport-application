package service

import (
	"airport/pkg/database"
	"airport/pkg/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAvailableSeatsByDestination(origin string, destination string) []model.Seat {
	var seats []model.Seat
	database.GetClient().Joins("inner join flights on flights.id = seats.flight_id").Where("flights.destination_id = ?", destination).Where("flights.origin_id = ?", origin).Where("seats.status = 'EMPTY'").Preload(clause.Associations).Find(&seats)
	if len(seats) == 0 {
		seats = getReservedSeatsByDestination(origin, destination)
	}

	return seats
}

func GetAllSeatsByDestination(origin string, destination string) []model.Seat {
	var seats []model.Seat
	database.GetClient().Joins("inner join flights on flights.id = seats.flight_id").Where("flights.destination_id = ?", destination).Where("flights.origin_id = ?", origin).Preload(clause.Associations).Find(&seats)

	return seats
}

func getReservedSeatsByDestination(origin string, destination string) []model.Seat {
	var seats []model.Seat
	database.GetClient().Joins("inner join flights on flights.id = seats.flight_id").Where("flights.destination_id = ?", destination).Where("flights.origin_id = ?", origin).Where("seats.status = 'RESERVED'").Preload(clause.Associations).Find(&seats)
	if len(seats) > 0 {
		seats = getExpiredReservationSeats(seats)
	}

	return seats
}

var getExpiredReservationSeats = func(seats []model.Seat) []model.Seat {
	ids := getSeatsIds(seats)
	var availableSeats []model.Seat
	database.GetClient().Joins("inner join sales on seats.id = sales.seat_id").Where("sales.seat_id in ?", ids).Where("TIMESTAMPDIFF(MINUTE, sales.reservation_date, now()) > 5").Preload(clause.Associations).Find(&availableSeats)

	return availableSeats
}

func UpdateExpiredReservedSeats(client *gorm.DB) {
	var foundSales []model.Sale
	client.Where("TIMESTAMPDIFF(MINUTE, reservation_date, now()) > 5").Find(&foundSales)
	var ids = getSeatsIdBySale(foundSales)
	client.Model(&model.Seat{}).Where("id in ?", ids).Update("status", "EMPTY")
}

func getSeatsIds(seats []model.Seat) []int {
	var ids []int
	for _, element := range seats {
		ids = append(ids, element.ID)
	}

	return ids
}

func getSeatsIdBySale(sales []model.Sale) []int {
	var ids []int
	for _, element := range sales {
		ids = append(ids, element.SeatID)
	}

	return ids
}
