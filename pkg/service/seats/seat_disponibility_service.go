package service

import (
	"strconv"
	"time"
)

type Route struct {
	origin      int
	destination int
}

var priceMap = make(map[Route]int)

var getSeatAvailability = func(origin int, destination int) int {
	var routeToFind = Route{origin: origin, destination: destination}
	value, exists := priceMap[routeToFind]
	if exists {
		return value
	}
	return loadAndMaintainAvailability(routeToFind)
}

func loadAndMaintainAvailability(routeToFind Route) int {
	var origin = strconv.Itoa(routeToFind.origin)
	var destination = strconv.Itoa(routeToFind.destination)
	availability := loadAvailability(routeToFind, origin, destination)
	channel := make(chan int)
	go updateAvailability(origin, destination, channel)
	go maintainAvailability(routeToFind, channel)

	return availability
}

// TODO: cada vez que se llama queda en un loop infinito, es necesario? y si solamente se llama bajo demanda?
func updateAvailability(origin string, destination string, channel chan int) {
	for {
		time.Sleep(1 * time.Second)
		channel <- CalculateAvailability(origin, destination)
	}
}

func maintainAvailability(routeToFind Route, channel chan int) {
	for update := range channel {
		priceMap[routeToFind] = update
	}
}

func loadAvailability(routeToFind Route, origin string, destination string) int {
	var availability = CalculateAvailability(origin, destination)
	priceMap[routeToFind] = availability

	return availability
}
