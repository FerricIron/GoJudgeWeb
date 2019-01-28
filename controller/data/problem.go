package data

import (
	"encoding/json"
	"github.com/ferriciron/GoJudgeWeb/common"
	"github.com/ferriciron/GoJudgeWeb/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetProblemInfo(c *gin.Context) {
	pid, err := strconv.Atoi(c.Param("pid"))
	if err != nil {
		model.WarnLog(c, err.Error())
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"errCode": common.InvalidParams,
			"message": "Invalid params",
		})
		return
	}
	ret, err := model.SelectProblem(pid)
	if err != nil {
		model.ErrLog(c, err.Error())
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"errCode": common.DataBaseUnavaliable,
			"message": "Database unavaliable",
		})
		return
	}
	retData,err:=json.Marshal(struct {
		ProblemId   int    	`json:"pid"`
		ProblemName string 	`json:"name"`
		Author      string 	`json:"author"`
		Description string 	`json:"description"`
		Property    string	`json:"property"`
		SubmitCount string 	`json:"submit"`
		Solved      string 	`json:"solved"`
		TimeLimit   string 	`json:"timelimit"`
		MemoryLimit string	`json:"memorylimit"`
	}{
		ProblemId:   ret.ProblemId,
		ProblemName: ret.ProblemName,
		Author:      ret.Author,
		Description: ret.Description,
		Property:	 ret.Property,
		SubmitCount: ret.SubmitCount,
		Solved:      ret.Solved,
		TimeLimit:   ret.TimeLimit,
		MemoryLimit: ret.MemoryLimit,
	})
	if err!=nil{
		model.ErrLog(c,err.Error())
		c.AbortWithStatusJSON(http.StatusOK,gin.H{
			"errCode":common.JSONComponentUnavaliable,
			"message":"Json marshal error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"errCode":common.Success,
		"message":retData,
	})
}
