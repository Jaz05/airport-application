package router

import (
	"airport/pkg/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Flights endpoints
	r.GET("/flights", controller.GetFlights)

	// Sales endpoints
	r.POST("/sales", controller.CreateSale)
	r.POST("/sales/:sale_id/payment", controller.CreatePayment)

	// Seats endpoints
	r.GET("/seats", controller.GetSeats)

	return r
}
