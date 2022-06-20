package controller

import (
	service "airport/pkg/service/seats"
	"github.com/gin-gonic/gin"
)

type seatResponse struct {
	Id       int
	Price    float32
	Position string
}

// GetSeats godoc
// @Summary  Get Seats by Origin and Destination
// @Tags     Seats
// @Produce  json
// @Param    origin       query    int  true  "Flight origin ID"
// @Param    destination  query    int  true  "Flight destination ID"
// @Success  200          {array}  seatResponse
// @Router   /seats [get]
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
