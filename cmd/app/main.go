package main

import (
	"airport/pkg/cron"
	"airport/pkg/database"
	"airport/pkg/loader"
	"airport/pkg/router"
)

func main() {
	r := router.SetupRouter()
	loader.LoadTables(database.GetClient())
	runCronTasks()
	err := r.Run(":8080")
	if err != nil {
		panic("Run error!")
	}
}

func runCronTasks() {
	go cron.UpdateExpiredReserved()
}
