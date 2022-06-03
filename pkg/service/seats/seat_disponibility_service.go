package service

import (
	"airport/pkg/database"
	"strconv"
	"time"
)

type Route struct {
	origin      int
	destination int
}

var priceMap = make(map[Route]int)

func GetSeatDisponibility(origin int, destination int) int {
	var routeToFind = Route{origin: origin, destination: destination}
	value, exists := priceMap[routeToFind]
	if exists {
		return value
	}
	return loadAndMaintainDisponiblity(routeToFind)
}

func loadAndMaintainDisponiblity(routeToFind Route) int {
	var origin = strconv.Itoa(routeToFind.origin)
	var destination = strconv.Itoa(routeToFind.destination)
	value := loadDisponibility(routeToFind, origin, destination)
	channel := make(chan int)
	go updateDisponibility(origin, destination, channel)
	go maintainDisponibility(routeToFind, channel)

	return value
}

func updateDisponibility(origin string, destination string, channel chan int) {
	for {
		time.Sleep(1 * time.Second)
		channel <- CalculateDisponibility(database.GetClient(), origin, destination)
	}
}

func maintainDisponibility(routeToFind Route, channel chan int) {
	for update := range channel {
		priceMap[routeToFind] = update
	}
}

func loadDisponibility(routeToFind Route, origin string, destination string) int {
	var disponibility = CalculateDisponibility(database.GetClient(), origin, destination)
	priceMap[routeToFind] = disponibility

	return disponibility
}
