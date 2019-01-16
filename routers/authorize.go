package routers

import (
	"github.com/ferriciron/GoJudgeWeb/controller/authorize"
	"github.com/gin-gonic/gin"
)

func setAuthorizeRouters(group *gin.RouterGroup) {
	group.POST("login", authorize.Login)
	group.POST("register", authorize.Register)
}
