package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
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

	cmd := exec.Command("rasdial", "sien", "/D")
	out, _ := cmd.CombinedOutput()
	fmt.Println(string(out))

	var computer WebComputer
	userFile := "C:/computer.json"
	fin, err := os.Open(userFile)
	defer fin.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	buf, err := ioutil.ReadFile(userFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	json.Unmarshal(buf, &computer)
	computer.Status = 3

	c, err := json.Marshal(computer)
	body := bytes.NewBuffer([]byte(c))
	res, err := http.Post("http://0.0.0.0:8081/v1/server/used?token=", "application/json;charset=utf-8", body)
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
	fmt.Println(result)
}
