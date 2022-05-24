package router

import (
	"airport/pkg/controller"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/flights", func(c *gin.Context) {
		c.JSON(200, controller.GetAllFlights())
	})

	return r
}
