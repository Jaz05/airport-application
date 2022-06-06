package service

import (
	"airport/pkg/database"
	"airport/pkg/loader"
	"airport/pkg/model"
	"airport/pkg/testutils"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestGetAllSeatsByDestinationShouldReturnAllSeats(t *testing.T) {
	testutils.BeforeEach()

	var flights = []model.Flight{{OriginID: 1, DestinationID: 2}, {OriginID: 3, DestinationID: 2}, {OriginID: 1, DestinationID: 3}}
	var seats = []model.Seat{{FlightID: 1, Status: "EMPTY"}, {FlightID: 1, Status: "EMPTY"}, {FlightID: 2, Status: "EMPTY"}, {FlightID: 1, Status: "RESERVED"}}
	db := database.GetInMemoryClient()
	loader.LoadTables(db)
	db.Create(flights)
	db.Create(seats)
	foundSeats := GetAllSeatsByDestination(db, "1", "2")
	if len(foundSeats) != 3 {
		t.Fail()
	}
}

func TestGetAllSeatsByDestinationShouldReturnReservedSeats(t *testing.T) {
	testutils.BeforeEach()

	currentTime := time.Now()
	expiredTime := currentTime.Add(-time.Minute * 10)
	var flights = []model.Flight{{OriginID: 1, DestinationID: 2, Date: currentTime}, {OriginID: 1, DestinationID: 3, Date: expiredTime}}
	var seats = []model.Seat{{FlightID: 1, Status: "OCCUPIED"}, {FlightID: 2, Status: "OCCUPIED"}, {FlightID: 1, Status: "RESERVED"}, {FlightID: 2, Status: "RESERVED"}}
	db := database.GetInMemoryClient()
	loader.LoadTables(db)
	db.Create(flights)
	db.Create(seats)
	getExpiredReservationSeats = func(seats []model.Seat, client *gorm.DB) []model.Seat {
		return []model.Seat{{FlightID: 2, Status: "RESERVED"}}
	}
	foundSeats := GetAllSeatsByDestination(db, "1", "2")
	if len(foundSeats) != 2 {
		t.Fail()
	}
}

func TestGetAvailableSeatsByDestinationWithNoAvailableSeatsShouldReturnEmpty(t *testing.T) {
	testutils.BeforeEach()

	currentTime := time.Now()
	expiredTime := currentTime.Add(-time.Minute * 10)
	var flights = []model.Flight{{OriginID: 1, DestinationID: 2, Date: currentTime}, {OriginID: 1, DestinationID: 3, Date: expiredTime}}
	var seats = []model.Seat{{FlightID: 1, Status: "OCCUPIED"}, {FlightID: 2, Status: "OCCUPIED"}, {FlightID: 1, Status: "OCCUPIED"}, {FlightID: 2, Status: "OCCUPIED"}}
	db := database.GetInMemoryClient()
	loader.LoadTables(db)
	db.Create(flights)
	db.Create(seats)
	foundSeats := GetAvailableSeatsByDestination(db, "1", "2")
	if len(foundSeats) != 0 {
		t.Fail()
	}
}
