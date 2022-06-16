package controller

import (
	"airport/pkg/model"
	"airport/pkg/service/sales"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type salesRequestBody struct {
	Sales []saleRequestBody `json:"sales"`
}

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

type salesResponseBody struct {
	Sales []saleResponseBody `json:"sales"`
	Token string             `json:"token"`
}

// CreateSales godoc
// @Summary      Creates a sale
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param body body salesRequestBody true "request body"
// @Success      200  {object}  salesResponseBody
// @Router       /sales [post]
func CreateSales(c *gin.Context) {
	var salesBody salesRequestBody
	if err := c.BindJSON(&salesBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var salesResponses []saleResponseBody
	token := uuid.New().String()

	for _, body := range salesBody.Sales {
		if _, err := sales.BookFlightSeat(body.SeatId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		var sale model.Sale
		sale, err := sales.SaveSale(body.SeatId, body.Dni, body.Name, body.Surname, token)
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

		salesResponses = append(salesResponses, saleResponse)
	}

	var response salesResponseBody
	response.Sales = salesResponses
	response.Token = token

	c.JSON(http.StatusOK, response)
}
