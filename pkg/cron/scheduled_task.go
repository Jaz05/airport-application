package cron

import (
	"airport/pkg/database"
	service "airport/pkg/service/seats"
	"time"
)

func UpdateExpiredReserved() {
	for {
		time.Sleep(10 * time.Second)
		service.UpdateExpiredReservedSeats(database.GetClient())
	}
}
