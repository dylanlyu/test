package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://0.0.0.0:8081/v1/account/read?token=d0U2U011VnRuaE52UGIvbmpiUEFkY2dzdU9yRGhiZVhYMjJNSm52YVlsSmJwaXNRcGV1eXFrYTRuOEg1K3BrS2M4MUdmaVVFWCtYVVRMOTkzMUtzQ2c9PS7vv70=")
	userSetting := "C:/PTT/tmp/user.json"
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fout, err := os.Create(userSetting)
	defer fout.Close()
	if err != nil {
		fmt.Println(userSetting, err)
		return
	}
	fout.Write([]byte(body))
}
