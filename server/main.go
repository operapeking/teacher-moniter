package main

import (
	"encoding/json"
	"flag"
	"time"

	"github.com/gin-gonic/gin"
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

	status := Status{
		Time:     time.Now(),
		Ip:       "127.0.0.1",
		IsComing: false,
		Who:      "none",
	}
	var send Send
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Service started")
	})
	r.GET("/query", func(ctx *gin.Context) {
		ctx.JSON(200, status)
	})
	r.POST("/modify", func(ctx *gin.Context) {
		json.NewDecoder(ctx.Request.Body).Decode(&send)
		status.Time = time.Now()
		status.Ip = ctx.ClientIP()
		status.IsComing = send.IsComing
		status.Who = send.Who
		ctx.String(200, "OK")
	})
	r.Run(ip + ":" + port)
}
