package controller

import (
	"airport/pkg/model"
	"airport/pkg/service"
	"github.com/gin-gonic/gin"
)

func GetFlights(c *gin.Context) {
	destination := c.Request.URL.Query().Get("destination")
	if len(destination) > 0 {
		c.JSON(200, getAllFlightsByDestination(destination))
	} else {
		c.JSON(200, getAllFlights())
	}
}

func getAllFlights() []model.Flight {
	return service.GetAllFlights()
}

func getAllFlightsByDestination(destination string) []model.Flight {
	return service.GetAllFlightsByDestination(destination)
}
