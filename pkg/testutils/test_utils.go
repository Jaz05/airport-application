package testutils

import (
	"airport/pkg/database"
	"airport/pkg/loader"
)

func BeforeEach() {
	db := database.GetInMemoryClient()
	db.Exec("DELETE FROM flights")
	db.Exec("DELETE FROM passengers")
	db.Exec("DELETE FROM sales")
	db.Exec("DELETE FROM seats")
	db.Exec("DELETE FROM seat_types")
	loader.LoadTables(db)
}

func MockGetSeatAvailability(availability int) func(origin int, destination int) int {
	return func(origin int, destination int) int {
		return availability
	}
}

func MockData(data interface{}) {
	db := database.GetInMemoryClient()
	db.Create(data)
}
