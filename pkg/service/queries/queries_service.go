package queries

import (
	"fmt"
	"io"
	"net/http"
)

// DelayGetUserInfo llamado a api con delay
func DelayGetUserInfo() (string, error) {
	const url = "https://reqres.in/api/users/2?delay=2"
	resp, err := http.Get(url)

	if err != nil {
		return "", fmt.Errorf("error fetching url: %q", url)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	response := string(respBody)
	return response, nil
}
