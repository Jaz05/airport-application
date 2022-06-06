package loader

import (
	"airport/pkg/model"
	"gorm.io/gorm"
)

func LoadTables(db *gorm.DB) {
	if err := db.AutoMigrate(&model.Airport{}, &model.Flight{}, &model.Passenger{}, &model.Place{}, &model.Sale{}, &model.SeatType{}, &model.Seat{}); err != nil {
		println(err.Error())
	}
}
