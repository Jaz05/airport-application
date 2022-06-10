package queries

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
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
