package service

import (
	"airport/pkg/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllSeatsByDestination(client *gorm.DB, origin string, destination string) []model.Seat {
	var seats []model.Seat
	client.Joins("inner join flights on flights.id = seats.flight_id").Where("flights.destination_id = ?", destination).Where("flights.origin_id = ?", origin).Where("seats.status = 'EMPTY'").Preload(clause.Associations).Find(&seats)
	if len(seats) == 0 {
		seats = getReservedSeatsByDestination(origin, destination, client)
	}

	return seats
}

func getReservedSeatsByDestination(origin string, destination string, client *gorm.DB) []model.Seat {
	var seats []model.Seat
	client.Joins("inner join flights on flights.id = seats.flight_id").Where("flights.destination_id = ?", destination).Where("flights.origin_id = ?", origin).Where("seats.status = 'RESERVED'").Preload(clause.Associations).Find(&seats)
	if len(seats) > 0 {
		seats = getExpiredReservationSeats(seats, client)
	}

	return seats
}

var getExpiredReservationSeats = func(seats []model.Seat, client *gorm.DB) []model.Seat {
	ids := getSeatsIds(seats)
	var availalbleSeats []model.Seat
	client.Joins("inner join sales on seats.id = sales.seat_id").Where("sales.seat_id in ?", ids).Where("TIMESTAMPDIFF(MINUTE, sales.reservation_date, now()) > 5").Preload(clause.Associations).Find(&availalbleSeats)

	return availalbleSeats
}

func getSeatsIds(seats []model.Seat) []int {
	var ids []int
	for _, element := range seats {
		ids = append(ids, element.ID)
	}

	return ids
}
