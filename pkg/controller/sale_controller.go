package controller

import (
	"airport/pkg/model"
	"airport/pkg/service/sales"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

type paymentResponseBody struct {
	Result string `json:"result"`
}

// CreateSale godoc
// @Summary      Creates a sale
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param body body saleRequestBody true "request body"
// @Success      200  {object}  saleResponseBody
// @Router       /sales [post]
func CreateSale(c *gin.Context) {
	var body saleRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

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

// CreatePayment godoc
// @Summary      Creates a payment
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        sale_id   path      int  true  "Sale ID"
// @Param body body paymentRequestBody true "request body"
// @Success      200  {object}  paymentResponseBody
// @Router       /sales/:sale_id/payment [post]
func CreatePayment(c *gin.Context) {
	var body paymentRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	sale, err := sales.GetSale(c.Param("sale_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = sales.ProcessPayment(sale, body.CardNumber, body.SecurityNumber, body.ExpirationDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, paymentResponseBody{
		Result: "Payment successful",
	})
}
