package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

var host string

func SendModify(status bool) {
	body, _ := json.Marshal(status)
	http.Post(host+"/modify", "application/json", bytes.NewBuffer(body))
}

func main() {
	args := os.Args
	host = args[1]
	if args[2] == "1" {
		SendModify(true)
	} else {
		SendModify(false)
	}
}
