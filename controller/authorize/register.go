package authorize

import (
	"github.com/ferriciron/GoJudgeWeb/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type registerForm struct {
	Username 		string 	`form:"username"`
	Password 		string 	`form:"password"`
	nickname 		string 	`form:"nickname"`
	description		string 	`form:"description"`
	sid				int		`form:"sid"`
}
func Register(c *gin.Context)  {
	var register registerForm
	if err:=c.ShouldBind(&register);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"errCode":common.InvalidForm,"message":err.Error()})
	}

	
}