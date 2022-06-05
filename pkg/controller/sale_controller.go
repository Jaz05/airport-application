package controller

import (
	"airport/pkg/model"
	"airport/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type SaleRequestBody struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Dni     int64  `json:"dni"`
	SeatId  int    `json:"seat_id"`
}

type SaleResponseBody struct {
	ID              int             `json:"id"`
	Passenger       model.Passenger `json:"passenger"`
	SeatID          int             `json:"seat_id"`
	Price           float32         `json:"price"`
	ReservationDate time.Time       `json:"reservation_date"`
}

type paymentRequestBody struct {
	CardNumber     int64  `json:"card_number"`
	SecurityNumber int    `json:"security_number"`
	ExpirationDate string `json:"expiration_date"`
}

func CreateSale(c *gin.Context) {
	var body SaleRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if _, err := service.BookFlightSeat(body.SeatId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var sale model.Sale
	sale, err := service.SaveSale(body.SeatId, body.Dni, body.Name, body.Surname)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var saleResponse SaleResponseBody
	saleResponse.ID = sale.ID
	saleResponse.Price = sale.Price
	saleResponse.Passenger = sale.Passenger
	saleResponse.ReservationDate = sale.ReservationDate
	saleResponse.SeatID = sale.SeatID

	c.JSON(http.StatusOK, saleResponse)
}

func CreatePayment(c *gin.Context) {
	//saleID := c.Param("name")
	var body SaleRequestBody

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

}
