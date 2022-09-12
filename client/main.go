package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var host string

var gone = 0

type Status struct {
	Time     time.Time `json:"time"`
	Ip       string    `json:"ip"`
	IsComing bool      `json:"isComing"`
}

func SendQuery() {
	var status Status
	resp, _ := http.Get(host + "/query")
	json.NewDecoder(resp.Body).Decode(&status)
	if status.IsComing && gone != 2 {
		fmt.Println("有人来了！来自：" + status.Ip + " 更新时间：" + status.Time.Format("15:04:05"))
		gone = 2
	} else if !status.IsComing && gone != 1 {
		fmt.Println("有人走了！来自：" + status.Ip + " 更新时间：" + status.Time.Format("15:04:05"))
		gone = 1
	}
}

func SendModify(status bool) {
	body, _ := json.Marshal(status)
	http.Post(host+"/modify", "application/json", bytes.NewBuffer(body))
}

func Query() {
	for {
		SendQuery()
		time.Sleep(time.Second)
	}
}

func Modify() {
	var input string
	for {
		fmt.Scan(&input)
		if input == "1" {
			SendModify(true)
		} else {
			SendModify(false)
		}
	}
}

func main() {
	fmt.Println("请输入服务器地址，如 http://192.168.1.1:10086")
	fmt.Scan(&host)
	go Query()
	Modify()
}
