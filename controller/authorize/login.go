package authorize

import (
	"github.com/ferriciron/GoJudgeWeb/common"
	"github.com/ferriciron/GoJudgeWeb/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required`
}

func CheckLoginForm(form loginForm) bool {
	if !(len(form.Username) > 6 && len(form.Username) < 20) {
		return false
	}
	if !(len(form.Password) > 8 && len(form.Password) < 20) {
		return false
	}
	for _, value := range form.Username {
		if !(value >= '0' && value <= '9' || value >= 'a' && value <= 'z' || value >= 'A' && value <= 'Z' || value != '_') {
			return false
		}
	}
	return true
}

func Login(c *gin.Context) {
	var login loginForm
	if err := c.ShouldBind(&login); err != nil {
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
				"errCode": common.InvalidForm,
				"message": err.Error(),
			})
		return
	}
	if !CheckLoginForm(login) {
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
				"errCode": common.InvalidForm,
				"message": "username or password not invalid",
			})
		return
	}
	uid, privilege, err := model.CheckUserPassword(login.Username, login.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
				"errCode": common.UserNotExist,
				"message": err.Error(),
			})
		return
	}
	var j common.JWT
	token, err := j.GenerateToken(uid, privilege)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"errCode": common.TokenComponentUnavaliable,
			"message": err.Error(),
		})
		return
	}
	c.Header("token", token)
	c.JSON(http.StatusOK, gin.H{
		"errCode": common.Success,
		"message": "ok",
	})
}
