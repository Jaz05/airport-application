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

// CreatePayment godoc
// @Summary  Creates a payment
// @Tags     Payments
// @Accept   json
// @Produce  json
// @Param    body  body  paymentRequestBody  true  "Request Body"
// @Success  201
// @Failure  400  {object}  map[string]string
// @Router   /payment [post]
func CreatePayment(c *gin.Context) {
	var body paymentRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

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

	c.Status(http.StatusCreated)
}
