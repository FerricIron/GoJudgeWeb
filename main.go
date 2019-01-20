package main

import (
	"fmt"
	"github.com/ferriciron/GoJudgeWeb/common"
	"github.com/ferriciron/GoJudgeWeb/routers"
	"github.com/gin-gonic/gin"
)
func main() {
	common.ParseConfig()
	go common.ConfigFileWatching()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	routers.SetRouters(r)
	r.Run(fmt.Sprintf("%s:%s",common.GlobalConfig.WebServer.Address,common.GlobalConfig.WebServer.Port))
}
