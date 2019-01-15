package routers

import (
	"github.com/ferriciron/GoJudgeWeb/controller/api"
	"github.com/gin-gonic/gin"
)

func SetApiRouters(group *gin.RouterGroup){
	group.Use(api.JWTAuth)

}
