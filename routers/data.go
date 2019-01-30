package routers

import (
	"github.com/ferriciron/GoJudgeWeb/controller/data"
	"github.com/ferriciron/GoJudgeWeb/controller/data/user"
	"github.com/gin-gonic/gin"
)

func SetDataRouters(group *gin.RouterGroup) {
	group.GET("/schools", data.GetSchoolList)
	group.GET("/problems", data.GetProblemsList)
	group.GET("/status",data.GetStatusList)
	group.GET("/problem/:pid", data.GetProblemInfo)

	group.GET("/user/:uid/email/confirm", user.UserEmailConfirm)
}
