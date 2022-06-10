package controller

import (
	"airport/pkg/model"
	"airport/pkg/service/queries"
	"airport/pkg/service/sales"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type saleRequestBody struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Dni     int64  `json:"dni"`
	SeatId  int    `json:"seat_id"`
}

type saleResponseBody struct {
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

type Result string

func CreateSale(c *gin.Context) {
	var body saleRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// llamados concurrentes a apis que tardan
	channel := make(chan string)
	go func() {
		var response, _ = queries.DelayGetUserInfo()
		channel <- response
	}()
	go func() {
		var response, _ = queries.DelayGetUserInfo()
		channel <- response
	}()
	go func() {
		var response, _ = queries.DelayGetUserInfo()
		channel <- response
	}()
	var responses []string
	for i := 0; i < 3; i++ {
		response := <-channel
		responses = append(responses, response)
	}
	fmt.Println(responses)

	if _, err := sales.BookFlightSeat(body.SeatId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var sale model.Sale
	sale, err := sales.SaveSale(body.SeatId, body.Dni, body.Name, body.Surname)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var saleResponse saleResponseBody
	saleResponse.ID = sale.ID
	saleResponse.Price = sale.Price
	saleResponse.Passenger = sale.Passenger
	saleResponse.ReservationDate = sale.ReservationDate
	saleResponse.SeatID = sale.SeatID

	c.JSON(http.StatusOK, saleResponse)
}

func CreatePayment(c *gin.Context) {
	var body paymentRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	sale, err := sales.GetSale(c.Param("sale_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	err = sales.ProcessPayment(sale, body.CardNumber, body.SecurityNumber, body.ExpirationDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}
