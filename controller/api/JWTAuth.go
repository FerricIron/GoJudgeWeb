package api

import (
	"github.com/ferriciron/GoJudgeWeb/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTAuth(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		c.JSON(http.StatusOK, gin.H{
			"errCode": common.TokenNotExist,
			"message": "Token not exist",
		})
		//c.Redirect(http.StatusMovedPermanently, "authorize/login")
		c.Abort()
	}

}
