package main

import (
	"github.com/ferriciron/GoJudgeWeb/routers"
	"github.com/gin-gonic/gin"
)
func main(){

	r:=gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	routers.SetRouters(r)
	r.Run()
}