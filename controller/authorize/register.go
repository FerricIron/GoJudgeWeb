package authorize

import (
	"github.com/ferriciron/GoJudgeWeb/common"
	"github.com/ferriciron/GoJudgeWeb/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type registerForm struct {
	Username    string `form:"username"`
	Password    string `form:"password"`
	nickname    string `form:"nickname"`
	description string `form:"description"`
	sid         int    `form:"sid"`
}

func Register(c *gin.Context) {
	var register registerForm
	if err := c.ShouldBind(&register); err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"errCode": common.InvalidForm,
				"message": err.Error(),
			})
		c.Abort()
	}
	user := model.User{
		Username:    register.Username,
		Password:    register.Password,
		Nickname:    register.nickname,
		Description: register.description,
		Sid:         register.sid,
	}
	err := model.AddUser(&user)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"errCode": common.UserExist,
				"message": err.Error(),
			})
		c.Abort()
	}
	c.JSON(http.StatusOK,
		gin.H{
			"errCode": common.Success,
			"message": "Create User Success",
		})
}
