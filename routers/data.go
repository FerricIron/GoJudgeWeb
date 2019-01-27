package routers

import (
	"github.com/ferriciron/GoJudgeWeb/controller/data"
	"github.com/gin-gonic/gin"
)

func SetDataRouters(group *gin.RouterGroup) {
	group.GET("/schools",data.GetSchoolList)
}
