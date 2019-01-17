package authorize

import (
	"github.com/ferriciron/GoJudgeWeb/common"
	"github.com/ferriciron/GoJudgeWeb/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type registerForm struct {
	Username    string `form:"username" binding:"required`
	Password    string `form:"password" binding:"required`
	nickname    string `form:"nickname" binding:"required`
	description string `form:"description"`
	sid         int    `form:"sid"`
}
func CheckRegisterForm(form registerForm) bool {
	if !CheckLoginForm(loginForm{form.Username,form.Password}){
		return false
	}
	return true
}

func Register(c *gin.Context) {
	var register registerForm
	if err := c.ShouldBind(&register); err != nil {
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
				"errCode": common.InvalidForm,
				"message": err.Error(),
			})
		return
	}
	if !CheckRegisterForm(register){
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
				"errCode": common.InvalidForm,
				"message": "invalid form",
			})
		return
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
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
				"errCode": common.UserExist,
				"message": err.Error(),
			})
		return
	}
	c.JSON(http.StatusOK,
		gin.H{
			"errCode": common.Success,
			"message": "ok",
		})
	return
}
