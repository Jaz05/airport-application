package main

import (
	"airport/pkg/router"
)

func main() {
	r := router.SetupRouter()
	//loader.LoadTables(database.GetClient())
	r.Run(":8080")
}
