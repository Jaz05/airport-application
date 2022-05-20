package main

import "airport/pkg/router"

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
