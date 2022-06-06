package service

import (
	"airport/pkg/model"
	"airport/pkg/testutils"
	"testing"
)

func TestGetAllFlightsShouldReturnAllFlights(t *testing.T) {
	testutils.BeforeEach()
	testutils.MockData([]model.Flight{{}, {}, {}})

	foundFlights := GetAllFlights()
	if len(foundFlights) != 3 {
		t.Fail()
	}

}

func TestGetAllFlightsByDestinationShouldReturnTwoFlights(t *testing.T) {
	testutils.BeforeEach()
	testutils.MockData([]model.Flight{{DestinationID: 1}, {DestinationID: 1}, {DestinationID: 2}})

	foundFlights := GetAllFlightsByDestination("1")
	if len(foundFlights) != 2 {
		t.Fail()
	}
}
