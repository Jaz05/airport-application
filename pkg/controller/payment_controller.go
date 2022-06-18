package controller

import (
	"airport/pkg/service/payment"
	"airport/pkg/service/sales"
	"net/http"

	"github.com/gin-gonic/gin"
)

type paymentRequestBody struct {
	CardNumber     int64  `json:"card_number"`
	SecurityNumber int    `json:"security_number"`
	ExpirationDate string `json:"expiration_date"`
	Token          string `json:"token"`
}

type paymentResponseBody struct {
	Result string `json:"result"`
}

// CreatePayment godoc
// @Summary      Creates a payment
// @Tags         payment
// @Accept       json
// @Produce      json
// @Param body body paymentRequestBody true "request body"
// @Success      200  {object}  paymentResponseBody
// @Router       /payment [post]
func CreatePayment(c *gin.Context) {
	var body paymentRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// TODO: esta bien identificar a las sales por token y no por id?
	// si una persona hacer varias reservas recibe distintos tokens
	salesList, err := sales.GetSalesByToken(body.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	for _, sale := range salesList {
		err = sales.ValidateSale(sale)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	}

	err = payment.ProcessPayment(body.CardNumber, body.SecurityNumber, body.ExpirationDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	for _, sale := range salesList {
		err = sales.FulfillSale(sale)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, paymentResponseBody{
		Result: "Payment successful",
	})
}
