package router

import (
	"airport/pkg/controller"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := r.Group("")

	{
		flights := v1.Group("/flights")
		{
			flights.GET("", controller.GetFlights)
		}

		sales := v1.Group("/sales")
		{
			sales.POST("", controller.CreateSales)
			sales.POST(":sale_id/payment", controller.CreatePayment)
		}

		seats := v1.Group("/seats")
		{
			seats.GET("", controller.GetSeats)
		}

	}

	return r
}
