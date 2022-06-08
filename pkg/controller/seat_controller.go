package controller

import (
	"airport/pkg/service/seats"
	"github.com/gin-gonic/gin"
)

type seatResponse struct {
	Id       int
	Price    float32
	Position string
}

func GetSeats(c *gin.Context) {
	origin := c.Request.URL.Query().Get("origin")
	destination := c.Request.URL.Query().Get("destination")
	c.JSON(200, getAllSeatsByDestination(origin, destination))
}

func getAllSeatsByDestination(origin string, destination string) []seatResponse {
	var seats = service.GetAllSeatsByDestination(origin, destination)
	var responseList []seatResponse
	for _, seat := range seats {
		var price = service.CalculateSeatPrice(service.GetSeatAvailability, seat)
		var element = seatResponse{Id: seat.ID, Position: seat.SeatLocation, Price: price}
		responseList = append(responseList, element)
	}

	return responseList
}
