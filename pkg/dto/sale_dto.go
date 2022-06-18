package dto

import (
	"airport/pkg/model"
	"time"
)

type SalesRequestBody struct {
	Sales []SaleRequestBody `json:"sales"`
}

type SaleRequestBody struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Dni     int64  `json:"dni"`
	SeatId  int    `json:"seat_id"`
}

type SaleResponseBody struct {
	ID              int             `json:"id"`
	Passenger       model.Passenger `json:"passenger"`
	SeatID          int             `json:"seat_id"`
	Price           float32         `json:"price"`
	ReservationDate time.Time       `json:"reservation_date"`
}

type SalesResponseBody struct {
	Sales []SaleResponseBody `json:"sales"`
	Token string             `json:"token"`
}
