package controller

import (
	"airport/pkg/dto"
	"airport/pkg/service/sales"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// CreateSales godoc
// @Summary  Create Sales
// @Tags     Sales
// @Accept   json
// @Produce  json
// @Param    body  body      dto.SalesRequestBody  true  "Request Body"
// @Success  200   {object}  dto.SalesResponseBody
// @Failure  400   {object}  map[string]string
// @Router   /sales [post]
func CreateSales(c *gin.Context) {
	var salesBody dto.SalesRequestBody
	if err := c.BindJSON(&salesBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	token := uuid.New().String()
	createdSales, err := sales.CreateSales(salesBody, token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var saleResponse dto.SaleResponseBody
	var saleResponses []dto.SaleResponseBody
	for _, sale := range createdSales {
		saleResponse.ID = sale.ID
		saleResponse.Price = sale.Price
		saleResponse.Passenger = sale.Passenger
		saleResponse.ReservationDate = sale.ReservationDate
		saleResponse.SeatID = sale.SeatID

		saleResponses = append(saleResponses, saleResponse)
	}

	var response dto.SalesResponseBody
	response.Sales = saleResponses
	response.Token = token

	c.JSON(http.StatusOK, response)
}

// GetSalesByToken godoc
// @Summary  Get Sales by Token
// @Tags     Sales
// @Produce  json
// @Param    token  query     string  true  "Sale Token"
// @Success  200    {object}  dto.SalesResponseBody
// @Failure  400    {object}  map[string]string
// @Router   /sales [get]
func GetSalesByToken(c *gin.Context) {
	token := c.Request.URL.Query().Get("token")
	salesList, err := sales.GetSalesByToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, salesList)
}
