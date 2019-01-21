package main

import (
	"fmt"
	"github.com/ferriciron/GoJudgeWeb/common"
	"github.com/ferriciron/GoJudgeWeb/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)
func main() {
	common.ParseConfig()
	go common.ConfigFileWatching()
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:true,
		AllowMethods:     []string{"PUT", "PATCH","POST","GET"},
		AllowHeaders:     []string{"Origin","token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	routers.SetRouters(r)
	r.Run(fmt.Sprintf("%s:%s",common.GlobalConfig.WebServer.Address,common.GlobalConfig.WebServer.Port))
}
