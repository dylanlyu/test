package main

import (
	"net/http"
	"time"
)

func main() {
	time.Sleep(20 * time.Second)
	_, err := http.Get("http://localhost:8080/v1/client/give")
	if err != nil {
		// handle error
	}

}
