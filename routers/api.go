package routers

import (
	"github.com/ferriciron/GoJudgeWeb/controller/api"
	"github.com/ferriciron/GoJudgeWeb/controller/api/submit"
	"github.com/gin-gonic/gin"
)

func SetApiRouters(group *gin.RouterGroup) {
	group.Use(api.JWTAuth)
	group.POST("submit",submit.Submit)
}
