package main

import (
	"airport/pkg/database"
	"airport/pkg/model"
	"airport/pkg/router"
)

func main() {
	r := router.SetupRouter()
	AutoMigrateDB()
	r.Run(":8080")
}

//Create tables automatically
func AutoMigrateDB() {
	db := database.GetClient()
	if err := db.AutoMigrate(&model.Airport{}, &model.Flight{}, &model.Passenger{}, &model.Place{}, &model.Sale{}, &model.SeatType{}, &model.Seat{}); err != nil {
		println(err)
	}
}
