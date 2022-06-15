package main

import (
	"airport/cmd/app/docs"
	"airport/pkg/cron"
	"airport/pkg/database"
	"airport/pkg/loader"
	"airport/pkg/router"
)

// @title           Airport Application
// @version         1.0
// @description     Airport Rest API

// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	r := router.SetupRouter()
	docs.SwaggerInfo.BasePath = "/api/v1"

	loader.LoadTables(database.GetClient())
	// runCronTasks()
	err := r.Run(":8080")
	if err != nil {
		panic("Run error!")
	}
}

func runCronTasks() {
	go cron.UpdateExpiredReserved()
}
