package router

import (
	"airport/pkg/controller"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/flights", func(c *gin.Context) {
		destination := c.Request.URL.Query().Get("destination")
		if len(destination) > 0 {
			c.JSON(200, controller.GetAllFlightsByDestination(destination))
		} else {
			c.JSON(200, controller.GetAllFlights())
		}
	})

	r.POST("/sales", controller.CreateSale)

	r.GET("/seats", func(c *gin.Context) {
		origin := c.Request.URL.Query().Get("origin")
		destination := c.Request.URL.Query().Get("destination")
		c.JSON(200, controller.GetAllSeatsByDestination(origin, destination))
	})

	return r
}
