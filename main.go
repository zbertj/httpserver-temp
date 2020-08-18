package main

import (
	"fmt"
	"httpserver-temp/config/jsonconf"
	"httpserver-temp/router"
	"os"

	log "github.com/sirupsen/logrus"
)

func initlog() {
	// 日志初始化
	file, err := os.OpenFile(jsonconf.Conf.Logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.SetLevel(log.DebugLevel)
}

func main() {
	// 读取配置
	err := jsonconf.InitConfig("./config/config.json")
	if err != nil {
		return
	}

	initlog()

	// gin.SetMode(gin.ReleaseMode)
	r := router.InitRouter()
	r.Run(jsonconf.Conf.Addr)
}

// CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o camera-server main.go
// nohup ./camera-server 1>/dev/null &
