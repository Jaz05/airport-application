package main

import (
	"airport/pkg/database"
	"airport/pkg/loader"
	"airport/pkg/router"
)

func main() {
	r := router.SetupRouter()
	loader.LoadTables(database.GetClient())
	r.Run(":8080")
}
