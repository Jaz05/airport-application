package testutils

import "airport/pkg/database"

func BeforeEach() {
	db := database.GetInMemoryClient()
	db.Exec("DELETE FROM flights")
	db.Exec("DELETE FROM passengers")
	db.Exec("DELETE FROM sales")
	db.Exec("DELETE FROM seats")
}
