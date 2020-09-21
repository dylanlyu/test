package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
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

type UserComputer struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		var userComputer []UserComputer
		log.Println("請輸入郵遞區號或是輸入l查詢:")
		data, _, _ := reader.ReadLine()
		command := string(data)
		if command == "l" {
			res, err := http.Get("http://0.0.0.0:8081/v1/server/list?token=d0U2U011VnRuaE52UGIvbmpiUEFkY2dzdU9yRGhiZVhYMjJNSm52YVlsSmJwaXNRcGV1eXFrYTRuOEg1K3BrS2M4MUdmaVVFWCtYVVRMOTkzMUtzQ2c9PS7vv70=")
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
			//log.Println(string(result))
			json.Unmarshal(result, &userComputer)
			for i := 0; i < len(userComputer); i++ {
				fmt.Println("郵遞區號: " + userComputer[i].Code + "  城市: " + userComputer[i].Name)
			}

		} else {
			_, error := strconv.Atoi(command)
			if error != nil {
				log.Println("輸入錯誤!請重新再輸入一次!")
				continue
			}

			var computer WebComputer
			computer.Code = command

			c, err := json.Marshal(computer)
			body := bytes.NewBuffer([]byte(c))
			res, err := http.Post("http://0.0.0.0:8081/v1/server/userpickup?token=d0U2U011VnRuaE52UGIvbmpiUEFkY2dzdU9yRGhiZVhYMjJNSm52YVlsSmJwaXNRcGV1eXFrYTRuOEg1K3BrS2M4MUdmaVVFWCtYVVRMOTkzMUtzQ2c9PS7vv70=", "application/json;charset=utf-8", body)
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

			//log.Println(string(c))
			//log.Println(result)
			json.Unmarshal(result, &computer)

			if computer.Id > 0 {
				userOut := "C:/usercomputer.json"
				fout, err := os.Create(userOut)
				defer fout.Close()
				if err != nil {
					log.Println(userOut, err)
					return
				}
				fout.Write([]byte(result))
				//寫入computer.json

				ip := "/phone:" + computer.Ip
				cmd := exec.Command("rasdial", "", "", "", ip)
				out, _ := cmd.CombinedOutput()
				log.Println(string(out))
				break
			} else {
				log.Println("輸入位子沒有空閑的電腦或是沒有建立")
				continue
			}

		}
	}
}
