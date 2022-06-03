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
	r.Run(":8080")
}

func runCronTasks() {
	go cron.UpdateExpiredReserved()
}
