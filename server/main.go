package main

import (
	"encoding/json"
	"flag"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

type Status struct {
	Time     time.Time `json:"time"`
	Ip       string    `json:"ip"`
	IsComing bool      `json:"isComing"`
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
	}
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Service started")
	})
	r.GET("/query", func(ctx *gin.Context) {
		ctx.JSON(200, status)
	})
	r.POST("/modify", func(ctx *gin.Context) {
		status.Time = time.Now()
		status.Ip = ctx.ClientIP()
		body, _ := io.ReadAll(ctx.Request.Body)
		json.Unmarshal(body, &status.IsComing)
		ctx.String(200, "OK")
	})
	r.Run(ip + ":" + port)
}
