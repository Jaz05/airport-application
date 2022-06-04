package controller

import (
	"airport/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type saleRequestBody struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Dni     int64  `json:"dni"`
	SeatId  int    `json:"seat_id"`
}

func CreateSale(c *gin.Context) {
	var body saleRequestBody
	if err := c.BindJSON(&body); err != nil {
		return
	}

	if err := service.BookFlightSeat(body.SeatId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := service.SaveSale(body.SeatId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
