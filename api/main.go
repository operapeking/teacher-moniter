package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"net/http"
)

var host string

func SendModify(status bool) {
	body, _ := json.Marshal(status)
	http.Post(host+"/modify", "application/json", bytes.NewBuffer(body))
}

func main() {
	var ip string
	var port string
	var op string
	flag.StringVar(&ip, "ip", "127.0.0.1", "server ip")
	flag.StringVar(&port, "port", "10086", "server port")
	flag.StringVar(&op, "op", "0", "1 is coming, others are gone")
	flag.Parse()
	host = "http://" + ip + ":" + port
	println()
	if op == "1" {
		SendModify(true)
	} else {
		SendModify(false)
	}
}
