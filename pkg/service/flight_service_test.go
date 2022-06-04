package service

import (
	"airport/pkg/database"
	"airport/pkg/loader"
	"airport/pkg/model"
	"testing"
)

func TestGetAllFlightsShouldReturnAllFlights(t *testing.T) {
	var flights = []model.Flight{{}, {}, {}}
	db := database.GetInMemoryClient()
	loader.LoadTables(db)
	db.Create(flights)
	foundFlights := GetAllFlights()
	if len(foundFlights) != 3 {
		t.Fail()
	}

}

func TestGetAllFlightsByDestinationShouldReturnTwoFlights(t *testing.T) {
	var flights = []model.Flight{{DestinationID: 1}, {DestinationID: 1}, {DestinationID: 2}}
	db := database.GetInMemoryClient()
	loader.LoadTables(db)
	db.Create(flights)
	foundFlights := GetAllFlightsByDestination("1")
	if len(foundFlights) != 2 {
		t.Fail()
	}

}
