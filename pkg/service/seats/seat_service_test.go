package service

import (
	"airport/pkg/model"
	"airport/pkg/testutils"
	"testing"
	"time"
)

func TestGetAllSeatsByDestinationShouldReturnAllSeats(t *testing.T) {
	testutils.BeforeEach()
	testutils.MockData([]model.Flight{{OriginID: 1, DestinationID: 2}, {OriginID: 3, DestinationID: 2}, {OriginID: 1, DestinationID: 3}})
	testutils.MockData([]model.Seat{{FlightID: 1, Status: "EMPTY"}, {FlightID: 1, Status: "EMPTY"}, {FlightID: 2, Status: "EMPTY"}, {FlightID: 1, Status: "RESERVED"}})

	foundSeats := GetAllSeatsByDestination("1", "2")
	if len(foundSeats) != 3 {
		t.Fail()
	}
}

func TestGetAllSeatsByDestinationShouldReturnReservedSeats(t *testing.T) {
	testutils.BeforeEach()
	currentTime := time.Now()
	expiredTime := currentTime.Add(-time.Minute * 10)
	testutils.MockData([]model.Flight{{OriginID: 1, DestinationID: 2, Date: currentTime}, {OriginID: 1, DestinationID: 3, Date: expiredTime}})
	testutils.MockData([]model.Seat{{FlightID: 1, Status: "OCCUPIED"}, {FlightID: 2, Status: "OCCUPIED"}, {FlightID: 1, Status: "RESERVED"}, {FlightID: 2, Status: "RESERVED"}})

	getExpiredReservationSeats = func(seats []model.Seat) []model.Seat {
		return []model.Seat{{FlightID: 2, Status: "RESERVED"}}
	}
	foundSeats := GetAllSeatsByDestination("1", "2")
	if len(foundSeats) != 2 {
		t.Fail()
	}
}

func TestGetAvailableSeatsByDestinationWithNoAvailableSeatsShouldReturnEmpty(t *testing.T) {
	testutils.BeforeEach()
	currentTime := time.Now()
	expiredTime := currentTime.Add(-time.Minute * 10)
	testutils.MockData([]model.Flight{{OriginID: 1, DestinationID: 2, Date: currentTime}, {OriginID: 1, DestinationID: 3, Date: expiredTime}})
	testutils.MockData([]model.Seat{{FlightID: 1, Status: "OCCUPIED"}, {FlightID: 2, Status: "OCCUPIED"}, {FlightID: 1, Status: "OCCUPIED"}, {FlightID: 2, Status: "OCCUPIED"}})

	foundSeats := GetAvailableSeatsByDestination("1", "2")
	if len(foundSeats) != 0 {
		t.Fail()
	}
}
