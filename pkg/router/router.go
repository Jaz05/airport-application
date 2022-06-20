package router

import (
	"airport/pkg/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	v1 := r.Group("")
	{
		flights := v1.Group("/flights")
		{
			flights.GET("", controller.GetFlights)
		}

		seats := v1.Group("/seats")
		{
			seats.GET("", controller.GetSeats)
		}

		sales := v1.Group("/sales")
		{
			sales.POST("", controller.CreateSales)
			sales.GET("", controller.GetSalesByToken)
		}

		payment := v1.Group("/payment")
		{
			payment.POST("", controller.CreatePayment)
		}
	}

	return r
}
