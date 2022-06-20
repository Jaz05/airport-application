package queries

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type Result struct {
	code int
	msg  string
}

type Fetch func() Result

// FakeFetch returns a fetch func with url set via parameter
func FakeFetch(url string) Fetch {
	return func() Result {
		delay := rand.Intn(2) + 1

		// se pisa el url por parametro dado que es de mentira
		url = fmt.Sprintf("https://reqres.in/api/users/2?delay=%d", delay)

		resp, err := http.Get(url)
		if err != nil {
			return Result{
				code: 500,
				msg:  "",
			}
		}

		defer resp.Body.Close()
		respBody, err := io.ReadAll(resp.Body)
		response := string(respBody)
		return Result{200, response}
	}
}

// First (replicas[0]=queries.Fetch(url), ...)
// FanIn Concurrency Pattern, again
func first(replicas ...Fetch) Result {
	c := make(chan Result)
	fetchReplica := func(i int) { c <- replicas[i]() }
	for i := range replicas {
		go fetchReplica(i)
	}

	// devuelvo la respuesta de la replica mas rapida
	return <-c
}

// TODO: manejo de errores dentro de una goroutine?
// TODO: nice to have, pasar por parametro cuantas replicas se hacen

func FanInFetch(queries ...Fetch) ([]Result, error) {
	// varios llamados concurrentes a apis que tardan un tiempo variable usando goroutines,
	// me quedo con la respuesta mas rapida de cada fetch lanzando varios fetchs iguales con mas goroutines
	channel := make(chan Result)

	fetchFirst := func(i int) { channel <- first(queries[i], queries[i], queries[i]) }
	for i := range queries {
		i := i
		go fetchFirst(i)
	}

	var responses []Result
	timeout := time.After(3000 * time.Millisecond)

	for i := 0; i < len(queries); i++ {
		select {
		case response := <-channel:
			responses = append(responses, response)
		case <-timeout:
			err := errors.New("TIMEOUT")
			return nil, err
		}
	}

	return responses, nil
}
