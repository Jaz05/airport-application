package service

import (
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Route struct {
	origin      int
	destination int
}

var priceMap = make(map[Route]int)

var getSeatDisponibility = func(client *gorm.DB, origin int, destination int) int {
	var routeToFind = Route{origin: origin, destination: destination}
	value, exists := priceMap[routeToFind]
	if exists {
		return value
	}
	return loadAndMaintainDisponiblity(client, routeToFind)
}

func loadAndMaintainDisponiblity(client *gorm.DB, routeToFind Route) int {
	var origin = strconv.Itoa(routeToFind.origin)
	var destination = strconv.Itoa(routeToFind.destination)
	disponibility := loadDisponibility(client, routeToFind, origin, destination)
	channel := make(chan int)
	go updateDisponibility(client, origin, destination, channel)
	go maintainDisponibility(routeToFind, channel)

	return disponibility
}

// TODO: cada vez que se llama queda en un loop infinito, es necesario? y si solamente se llama bajo demanda?
func updateDisponibility(client *gorm.DB, origin string, destination string, channel chan int) {
	for {
		time.Sleep(1 * time.Second)
		channel <- CalculateDisponibility(client, origin, destination)
	}
}

func maintainDisponibility(routeToFind Route, channel chan int) {
	for update := range channel {
		priceMap[routeToFind] = update
	}
}

func loadDisponibility(client *gorm.DB, routeToFind Route, origin string, destination string) int {
	var disponibility = CalculateDisponibility(client, origin, destination)
	priceMap[routeToFind] = disponibility

	return disponibility
}
