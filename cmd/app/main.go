package main

import (
	"airport/pkg/database"
	"airport/pkg/model"
	"airport/pkg/router"
)

func main() {
	r := router.SetupRouter()
	autoMigrateDB()
	r.Run(":8080")
}

//Create tables automatically
func autoMigrateDB() {
	db := database.GetClient()
	if err := db.AutoMigrate(&model.Airport{}, &model.Flight{}, &model.Passenger{}, &model.Place{}, &model.Sale{}, &model.SeatType{}, &model.Seat{}); err != nil {
		println(err.Error())
	}
}
