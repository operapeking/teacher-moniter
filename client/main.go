package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var host string

type Status struct {
	Time     time.Time `json:"time"`
	Ip       string    `json:"ip"`
	IsComing bool      `json:"isComing"`
	Who      string    `json:"Who"`
}

type Send struct {
	IsComing bool   `json:"isComing"`
	Who      string `json:"who"`
}

var now, last Status

func SendQuery() {
	resp, _ := http.Get(host + "/query")
	json.NewDecoder(resp.Body).Decode(&now)
	if now.IsComing == last.IsComing && now.Who == last.Who {
		return
	}
	if now.IsComing {
		fmt.Println(now.Who + " 来了！来自：" + now.Ip + " 更新时间：" + now.Time.Format("15:04:05"))
	} else {
		fmt.Println(now.Who + " 走了！来自：" + now.Ip + " 更新时间：" + now.Time.Format("15:04:05"))
	}
	last = now
}

func SendModify(send Send) {
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(send)
	http.Post(host+"/modify", "application/json", body)
}

func Query() {
	for {
		SendQuery()
		time.Sleep(time.Second)
	}
}

func Modify() {
	var send Send
	for {
		fmt.Scan(&send.IsComing, &send.Who)
		SendModify(send)
	}
}

func main() {
	fmt.Println("请输入服务器地址，如 http://192.168.1.1:10086")
	fmt.Scan(&host)
	go Query()
	Modify()
}
