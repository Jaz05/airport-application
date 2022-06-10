package controller

import (
	"airport/pkg/model"
	"airport/pkg/service/queries"
	"airport/pkg/service/sales"
	"errors"
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

// First (replicas[0]=queries.DelayGetUserInfo, ...)
// FanIn Concurrency Pattern, again
func First(replicas ...func() string) string {
	c := make(chan string)
	fetchReplica := func(i int) { c <- replicas[i]() }
	for i := range replicas {
		go fetchReplica(i)
	}

	// devuelvo la respuesta de la replica mas rapida
	return <-c
}

func CreateSale(c *gin.Context) {
	var body saleRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// varios llamados concurrentes a apis que tardan un tiempo variable usando goroutines,
	// me quedo con la respuesta mas rapida de cada fetch lanzando varios fetchs iguales con mas goroutines
	channel := make(chan string)

	// FanIn Concurrency Pattern
	go func() {
		channel <- First(queries.DelayGetUserInfo, queries.DelayGetUserInfo, queries.DelayGetUserInfo)
	}()
	go func() {
		channel <- First(queries.DelayGetClimateInfo, queries.DelayGetClimateInfo, queries.DelayGetClimateInfo)
	}()
	go func() {
		channel <- First(queries.DelayGetDollarInfo, queries.DelayGetDollarInfo, queries.DelayGetDollarInfo)
	}()

	var responses []string
	timeout := time.After(3000 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case response := <-channel:
			responses = append(responses, response)
		case <-timeout:
			err := errors.New("TIMEOUT")
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	}
	//

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
