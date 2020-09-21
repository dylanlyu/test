package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type ClientAccount struct {
	Id         int    `json:"id"`
	Account    string `json:"account"`
	Pass       string `json:"pass"`
	IsError    int    `json:"iserror"`
	IsAuto     int    `json:"isauto"`
	IsUserSet  int    `json:"isuserset"`
	Gas        int    `json:"gas"`
	Action     int    `json:"action"`
	UpdateTime int64  `json:"updatetime"`
	Primary    string `json:"primary"`
	Secondary  string `json:"secondary"`
	Tertiary   string `json:"tertiary"`
}

func main() {
	var client ClientAccount

	userFile := "C:/user.json"
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
	json.Unmarshal(buf, &client)
	client.UpdateTime = time.Now().Unix()

	c, err := json.Marshal(client)
	body := bytes.NewBuffer([]byte(c))
	res, err := http.Post("http://0.0.0.0:8081/v1/account/update?token=d0U2U011VnRuaE52UGIvbmpiUEFkY2dzdU9yRGhiZVhYMjJNSm52YVlsSmJwaXNRcGV1eXFrYTRuOEg1K3BrS2M4MUdmaVVFWCtYVVRMOTkzMUtzQ2c9PS7vv70=", "application/json;charset=utf-8", body)
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
