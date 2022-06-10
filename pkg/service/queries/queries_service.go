package queries

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

// DelayGetUserInfo llamado a api con delay
// TODO: pasar por parametro el delay?
func DelayGetUserInfo() string {
	// tarda entre 1 y 3 segundos
	delay := rand.Intn(2) + 1
	url := fmt.Sprintf("https://reqres.in/api/users/2?delay=%d", delay)

	resp, err := http.Get(url)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	response := string(respBody)
	return response
}

func DelayGetClimateInfo() string {
	delay := rand.Intn(2) + 1
	url := fmt.Sprintf("https://reqres.in/api/users/2?delay=%d", delay)

	resp, err := http.Get(url)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	response := string(respBody)
	return response
}

func DelayGetDollarInfo() string {
	delay := rand.Intn(2) + 1
	url := fmt.Sprintf("https://reqres.in/api/users/2?delay=%d", delay)

	resp, err := http.Get(url)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	response := string(respBody)
	return response
}

type Result string

// First (replicas[0]=queries.DelayGetUserInfo, ...)
// FanIn Concurrency Pattern, again
func first(replicas ...func() string) string {
	c := make(chan string)
	fetchReplica := func(i int) { c <- replicas[i]() }
	for i := range replicas {
		go fetchReplica(i)
	}

	// devuelvo la respuesta de la replica mas rapida
	return <-c
}

// TODO: deshardcodear que sean si o si 3 queries
// TODO: nice to have, pasar por parametro cuantas replicas se hacen
func FanInFetch(queries ...func() string) ([]string, error) {
	// varios llamados concurrentes a apis que tardan un tiempo variable usando goroutines,
	// me quedo con la respuesta mas rapida de cada fetch lanzando varios fetchs iguales con mas goroutines
	channel := make(chan string)

	// FanIn Concurrency Pattern
	go func() {
		channel <- first(queries[0], queries[0], queries[0])
	}()
	go func() {
		channel <- first(queries[1], queries[1], queries[1])
	}()
	go func() {
		channel <- first(queries[2], queries[2], queries[2])
	}()

	var responses []string
	timeout := time.After(3000 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case response := <-channel:
			responses = append(responses, response)
		case <-timeout:
			err := errors.New("TIMEOUT")

			return nil, err
		}
	}
	//
	return responses, nil
}
