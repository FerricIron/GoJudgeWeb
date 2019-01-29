package user

import (
	"fmt"
	"github.com/ferriciron/GoJudgeWeb/common"
	"github.com/ferriciron/GoJudgeWeb/email"
	"github.com/ferriciron/GoJudgeWeb/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)
const expiredTime = 30  //minute
type confirmStruct struct {
	SendTime  	int64 	`form:"t" binding:"required"`
	Content		string	`form:"c" binding:"required"`
	Email 		string	`form:"e" binding:"required"`
}

func UserEmailConfirm(c *gin.Context){
	var confirm confirmStruct
	uid,err:=strconv.Atoi(c.Param("uid"))
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusOK,gin.H{
			"errCode":common.InvalidForm,
			"message":"uid invalid",
		})
		return
	}
	if err:=c.ShouldBind(&confirm);err!=nil{
		c.AbortWithStatusJSON(http.StatusOK,gin.H{
			"errCode":common.InvalidParams,
			"message":"InvalidForm",
		})
		return
	}
	user,err:=model.SelectUser(uid)
	if err!=nil{
		model.ErrLog(c,err.Error())
		c.AbortWithStatusJSON(http.StatusOK,gin.H{
			"errCode":common.DataBaseUnavaliable,
			"message":"Database unavaliable",
		})
		return
	}
	content:=email.GenerateContent(user.Username,user.Password,confirm.Email,confirm.SendTime)
	expired:=time.Unix(confirm.SendTime,0).Add(expiredTime*time.Minute).Unix()
	if strings.Compare(confirm.Content,content)!=0{
		fmt.Println(confirm.Content)
		fmt.Println(content)
		c.AbortWithStatusJSON(http.StatusOK,gin.H{
			"errCode":common.InvalidParams,
			"message":"Invalid content",
		})
		return
	}
	if time.Now().Unix()>expired{
		c.AbortWithStatusJSON(http.StatusOK,gin.H{
			"errCode":common.InvalidParams,
			"message":"The link is expired,please retry!",
		})
		return
	}
	user.Email=confirm.Email
	_,err=model.UpdateUser(user.Uid,user)
	if err!=nil{
		model.ErrLog(c,err.Error())
		c.AbortWithStatusJSON(http.StatusOK,gin.H{
			"errCode":common.DataBaseUnavaliable,
			"message":"Database unavaliable",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"errCode":common.Success,
		"message":"Update email success",
	})
}
