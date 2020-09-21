package main

import (
	"net/http"
)

func main() {
	_, err := http.Get("http://localhost:8080/v1/client/check")
	if err != nil {
		// handle error
	}

}
