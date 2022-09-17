package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type Send struct {
	IsComing bool   `json:"isComing"`
	Who      string `json:"who"`
}

var host string
var send Send

func SendModify() {
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(send)
	http.Post(host+"/modify", "application/json", body)
}

func main() {
	var ip, port, op string
	flag.StringVar(&ip, "ip", "127.0.0.1", "server ip")
	flag.StringVar(&port, "port", "10086", "server port")
	flag.StringVar(&op, "op", "0", "1 is coming, 0 is gone")
	flag.StringVar(&send.Who, "who", "someone", "who is coming")
	flag.Parse()
	host = "http://" + ip + ":" + port
	if op == "1" {
		send.IsComing = true
	} else if op == "0" {
		send.IsComing = false
	} else {
		fmt.Println("op 参数输入错误")
		return
	}
	SendModify()
}
