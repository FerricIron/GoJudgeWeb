package authorize

import (
	"github.com/ferriciron/GoJudgeWeb/common"
	"github.com/ferriciron/GoJudgeWeb/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func Login(c *gin.Context) {
	var login loginForm
	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"errCode": common.InvalidForm,
				"message": err.Error(),
			})
		c.Abort()
	}
	uid, privilege, err := model.CheckUserPassword(login.Username, login.Password)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"errCode": common.UserNotExist,
				"message": err.Error(),
			})
		c.Abort()
	}
	var j common.JWT
	token, err := j.GenerateToken(uid, privilege)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"errCode": common.TokenComponentUnavaliable,
			"message": err.Error(),
		})
		c.Abort()
	}
	c.Header("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, gin.H{
		"errCode": common.Success,
		"message": "ok",
	})
}
