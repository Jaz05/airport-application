package controller

import (
	"airport/pkg/model"
	"airport/pkg/service"
)

func GetAllFlights() []model.Flight {
	return service.GetAllFlights()
}
