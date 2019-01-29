package main

import (
	"fmt"
	"github.com/ferriciron/GoJudgeWeb/common"
	"github.com/ferriciron/GoJudgeWeb/email"
	"github.com/ferriciron/GoJudgeWeb/model"
	"github.com/ferriciron/GoJudgeWeb/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"syscall"
	"time"
)

func main() {
	email.TestSend()
	common.ParseConfig()
	go ConfigFileWatching()
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:     []string{"Origin", "token", "content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	routers.SetRouters(r)
	r.Run(fmt.Sprintf("%s:%s", common.GlobalConfig.WebServer.Address, common.GlobalConfig.WebServer.Port))
}
func ConfigFileWatching() {
	var lastChangeTime int64 = 0
	for {
		fileInfo, err := os.Stat(common.ConfigPath)
		if err != nil {
			time.Sleep(common.WatchingSleepTime * time.Second)
			continue
		}
		stat_t := fileInfo.Sys().(*syscall.Stat_t)
		changeTime := syscall.TimespecToNsec(stat_t.Mtim)
		if lastChangeTime == 0 {
			lastChangeTime = changeTime
		} else {
			if changeTime > lastChangeTime {
				lastChangeTime = changeTime
				common.ParseConfig()
				model.ServerLog("ConfigFile Changes")
			}
			time.Sleep(common.WatchingSleepTime * time.Second)
		}
	}
}
