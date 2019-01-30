package data

import (
	"encoding/json"
	"github.com/ferriciron/GoJudgeWeb/common"
	"github.com/ferriciron/GoJudgeWeb/model"
	"github.com/gin-gonic/gin"
	"net/http"
)
type statusQuery struct{
	Page int `form:"page" binding:"required"`
	Capacity int `form:"capacity" binding:"required"`
}
type statusLIstRetStruct struct{
	SubmitId   	int		`json:"id"`
	UserName   	string	`json:"username"`
	Uid        	int		`json:"uid"`
	Time       	int		`json:"time"`
	Language   	int		`json:"language"`
	ProblemName string	`json:"problemname"`
	ProblemId  	int		`json:"problemid"`
	Status     	int		`json:"status"`
	TimeCost  	int		`json:"timecost"`
}
func statusListRetWrapper(data []model.Submit,maxPage int)(interface{}){
	var retDt []statusLIstRetStruct
	for _, v := range data {
		retDt=append(retDt, statusLIstRetStruct{
			SubmitId:    v.SubmitId,
			UserName:    v.Username,
			Uid:         v.Uid,
			Time:        v.Time,
			Language:    v.Language,
			ProblemName: v.ProblemName,
			ProblemId:   v.ProblemId,
			Status:      v.Status,
			TimeCost:    v.TimeCost,
		})
	}
	return struct {
		MaxPage int 					`json:"maxPage"`
		Data	[]statusLIstRetStruct	`json:"data"`
	}{
		MaxPage: maxPage,
		Data:    retDt,
	}
}

func GetStatusList(c *gin.Context){
	var queryStruct statusQuery
	err:=c.ShouldBind(&queryStruct)
	if err!=nil{
		model.ServerLog(err.Error())
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
				"errCode":common.InvalidForm,
				"message":"InvalidForm",
			})
		return
	}
	data,maxPage,err:=model.SelectStatusList(queryStruct.Page,queryStruct.Capacity)
	if err!=nil{
		model.ServerLog(err.Error())
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
				"errCode":common.DataBaseUnavaliable,
				"message":"Database unavaliable",
			})
		return
	}
	ret,err:=json.Marshal(statusListRetWrapper(data,maxPage))
	if err!=nil{
		model.ServerLog(err.Error())
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
				"errCode":common.JSONComponentUnavaliable,
				"message":"Json Marshal error",
			})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"errCode":common.Success,
		"message":ret,
	})
}

