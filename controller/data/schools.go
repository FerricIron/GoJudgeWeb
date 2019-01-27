package data

import (
	"encoding/json"
	"github.com/ferriciron/GoJudgeWeb/common"
	"github.com/ferriciron/GoJudgeWeb/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSchoolList(c *gin.Context){
	list,err:=model.SelectAllSchool()
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
				"errCode":common.DataBaseUnavaliable,
				"message":"Database Unavaliable",
			})
		model.ErrLog(c,err.Error())
		return
	}
	ret,err:=json.Marshal(&list)
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
				"errCode":common.JSONComponentUnavaliable,
				"message":"Json marshal error",
			})
		model.ErrLog(c,err.Error())
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"errCode":0,
		"message":ret,
	})
}
