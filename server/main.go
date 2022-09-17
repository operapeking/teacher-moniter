package main

import (
	"encoding/json"
	"flag"
	"net/http"
	"strings"
	"time"
)

type Status struct {
	Time     time.Time `json:"time"`
	Ip       string    `json:"ip"`
	IsComing bool      `json:"isComing"`
	Who      string    `json:"who"`
}

type Send struct {
	IsComing bool   `json:"isComing"`
	Who      string `json:"who"`
}

func main() {
	var ip string
	var port string
	flag.StringVar(&ip, "ip", "127.0.0.1", "listening ip")
	flag.StringVar(&port, "port", "10086", "listening port")
	flag.Parse()
	http.HandleFunc("/query", QueryPage)
	http.HandleFunc("/modify", ModifyPage)
	http.ListenAndServe(ip+":"+port, nil)
}

var status Status
var send Send

func QueryPage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(status)
}

func ModifyPage(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&send)
	status.Time = time.Now()
	status.Ip = r.RemoteAddr[:strings.Index(r.RemoteAddr, ":")]
	status.IsComing = send.IsComing
	status.Who = send.Who
}
