package main

import (
	"bytes"
	"encoding/json"
	"flag"
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
	var ip, port string
	flag.StringVar(&ip, "ip", "127.0.0.1", "server ip")
	flag.StringVar(&port, "port", "10086", "server port")
	flag.BoolVar(&send.IsComing, "op", false, "1 is coming, 0 is gone")
	flag.StringVar(&send.Who, "who", "none", "who is coming")
	flag.Parse()
	host = "http://" + ip + ":" + port
	SendModify()
}
