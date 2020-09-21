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

type ClientAccount struct {
	Id         int    `json:"id"`
	Account    string `json:"account"`
	Pass       string `json:"pass"`
	IsError    int    `json:"iserror"`
	IsAuto     int    `json:"isauto"`
	IsUserSet  int    `json:"isuserset"`
	Gas        int    `json:"gas"`
	Action     int    `json:"action"`
	UpdateTime int    `json:"updatetime"`
	Primary    string `json:"primary"`
	Secondary  string `json:"secondary"`
	Tertiary   string `json:"tertiary"`
}

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
	var account ClientAccount
	var computer WebComputer
	//var result
	switchcity := 1
	userFile := "C:/user.json"
	userOut := "C:/computer.json"
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
	json.Unmarshal(buf, &account)
	//讀取user.json

	for {
		switch switchcity {
		case 1:
			computer.City = account.Primary
		case 2:
			computer.City = account.Secondary
		case 3:
			computer.City = account.Tertiary
		}

		c, err := json.Marshal(computer)
		body := bytes.NewBuffer([]byte(c))
		res, err := http.Post("http://0.0.0.0:8081/v1/server/pickup?token=d0U2U011VnRuaE52UGIvbmpiUEFkY2dzdU9yRGhiZVhYMjJNSm52YVlsSmJwaXNRcGV1eXFrYTRuOEg1K3BrS2M4MUdmaVVFWCtYVVRMOTkzMUtzQ2c9PS7vv70=", "application/json;charset=utf-8", body)
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
		//取得computer

		json.Unmarshal(result, &computer)

		if computer.Id > 0 {
			fout, err := os.Create(userOut)
			defer fout.Close()
			if err != nil {
				fmt.Println(userOut, err)
				return
			}
			fout.Write([]byte(result))
			//寫入computer.json
			break
		}
		switchcity++
		fmt.Println(switchcity)
		fmt.Println(computer)
		time.Sleep(5 * time.Second)
		if switchcity > 3 {
			switchcity = 1
		}

	}
	ip := "/phone:" + computer.Ip
	cmd := exec.Command("rasdial", "", "", "=", i)
	out, _ := cmd.CombinedOutput()
	fmt.Println(string(out))

}
