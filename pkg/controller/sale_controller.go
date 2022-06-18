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

	/*reps: how many sales are we doing*/
	reps := len(salesBody.Sales)
	/*reps: each sales returns a value through this channel*/
	channel := make(chan model.Seat)
	/*reps: if an error happens it returns an error, to this channel, instead */
	errors := make(chan error)
	// TODO: una alternativa es hacer un Optional/RightLeft del tipo SeatError{Seat, error} y mandar todo al mismo canal
	// TODO: mover todo esto a un service pq tiene logica de negocio el controller
	// TODO: si un solo asiento se reserva queda reservado y el otro no (implementar rollback)
	// TODO: validar que no tengan el mismo seat id
	for i := 0; i < reps; i++ {
		body := salesBody.Sales[i]
		go func() {
			seat, err := sales.BookFlightSeat(body.SeatId)
			if err != nil {
				errors <- err
			}
			channel <- seat
		}()
	}

	/*we expect an amount of answers equal to the length of sales*/
	/*we throw an error if at least one of them fails*/
	for i := 0; i < reps; i++ {
		select {
		case <-channel:
		case err := <-errors:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	}

	saleChannel := make(chan model.Sale)

	for i := 0; i < reps; i++ {
		body := salesBody.Sales[i]
		go func() {
			var sale model.Sale
			sale, err := sales.SaveSale(body.SeatId, body.Dni, body.Name, body.Surname, token)
			if err != nil {
				errors <- err
			}
			saleChannel <- sale

		}()
	}

	var saleResponse saleResponseBody
	for i := 0; i < reps; i++ {
		select {
		case sale := <-saleChannel:
			saleResponse.ID = sale.ID
			saleResponse.Price = sale.Price
			saleResponse.Passenger = sale.Passenger
			saleResponse.ReservationDate = sale.ReservationDate
			saleResponse.SeatID = sale.SeatID

			salesResponses = append(salesResponses, saleResponse)
		case err := <-errors:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	}

	//
	var response salesResponseBody
	response.Sales = salesResponses
	response.Token = token

	c.JSON(http.StatusOK, response)
}

// GetSalesByToken godoc
// @Summary      Get sales by token
// @Tags         sales
// @Produce      json
// @Param        token   query int  true  "Sale token"
// @Success      200  {array}  []model.Sale
// @Router       /sales [get]
func GetSalesByToken(c *gin.Context) {
	token := c.Request.URL.Query().Get("token")
	salesList, err := sales.GetSalesByToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, salesList)
}
