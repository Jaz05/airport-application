package router

import (
	"airport/pkg/controller"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

type saleRequestBody struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Dni     int64  `json:"dni"`
	SeatId  int64  `json:"seat_id"`
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/flights", func(c *gin.Context) {
		destination := c.Request.URL.Query().Get("destination")
		if len(destination) > 0 {
			c.JSON(200, controller.GetAllFlightsByDestination(destination))
		} else {
			c.JSON(200, controller.GetAllFlights())
		}
	})

	r.POST("/sales", func(c *gin.Context) {
		var newSaleRequestBody saleRequestBody
		if err := c.BindJSON(&newSaleRequestBody); err != nil {
			return
		}
		seatId := strconv.Itoa(int(newSaleRequestBody.SeatId))
		seat, err := controller.BookFlightSeat(seatId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, seat)
	})

	r.GET("/seats", func(c *gin.Context) {
		origin := c.Request.URL.Query().Get("origin")
		destination := c.Request.URL.Query().Get("destination")
		c.JSON(200, controller.GetAllSeatsByDestination(origin, destination))
	})

	return r
}
