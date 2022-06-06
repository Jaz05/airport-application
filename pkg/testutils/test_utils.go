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
	loader.LoadTables(db)
}

func MockData(data interface{}) {
	db := database.GetInMemoryClient()
	db.Create(data)
}
