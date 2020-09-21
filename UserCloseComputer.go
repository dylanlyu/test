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

	cmd := exec.Command("rasdial", "", "/D")
	out, _ := cmd.CombinedOutput()
	log.Println(string(out))

	var computer WebComputer
	userFile := "C:/usercomputer.json"
	fin, err := os.Open(userFile)
	defer fin.Close()
	if err != nil {
		log.Println(userFile, err)
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
	res, err := http.Post("http://0.0.0.0:8081/v1/server/used?token=d0U2U011VnRuaE52UGIvbmpiUEFkY2dzdU9yRGhiZVhYMjJNSm52YVlsSmJwaXNRcGV1eXFrYTRuOEg1K3BrS2M4MUdmaVVFWCtYVVRMOTkzMUtzQ2c9PS7vv70=", "application/json;charset=utf-8", body)
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
	log.Println(result)
}
