package router

import (
	"airport/pkg/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/flights", controller.GetFlights)

	r.POST("/sales", controller.CreateSale)

	r.GET("/seats", controller.GetSeats)

	return r
}
