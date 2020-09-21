package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type WebComputer struct {
	Id       int64     `json:"id"`
	Computer string    `json:"computer"`
	City     string    `json:"city"`
	Code     string    `json:"code"`
	Ip       string    `json:"ip"`
	Time     time.Time `json:"time"`
	Status   int       `json:"status"`
}

func main() {
	//var computer WebComputer
	//computer.Computer = "KHH-813-01"
	//c, err := json.Marshal(computer)
	//body := bytes.NewBuffer([]byte(c))
	res, err := http.Get("http://0.0.0.08081/")
	if err != nil {
		log.Fatal(err)
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(string(result))
}
