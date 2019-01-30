package data

import (
	"encoding/json"
	"github.com/ferriciron/GoJudgeWeb/common"
	"github.com/ferriciron/GoJudgeWeb/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type problemQuery struct {
	Page     int `form:"page" binding:"required"`
	Capacity int `form:"capacity" binding:"required"`
}
type problemListRetStruct struct {
	ProblemId   int    `json:"pid"`
	ProblemName string `json:"name"`
	Privilege   int    `json:"privilege"`
	SubmitCount string `json:"submit"`
	Solved      string `json:"solved"`
}

func problemListRetWrapper(data []model.Problem, maxPage int) interface{} {
	var retDt []problemListRetStruct
	for _, v := range data {
		retDt = append(retDt, problemListRetStruct{
			ProblemId:   v.ProblemId,
			ProblemName: v.ProblemName,
			Privilege:   v.Privilege,
			SubmitCount: v.SubmitCount,
			Solved:      v.Solved,
		})
	}
	return struct {
		MaxPage int                    `json:"maxPage"`
		Data    []problemListRetStruct `json:"data"`
	}{
		MaxPage: maxPage,
		Data:    retDt,
	}
}

func GetProblemsList(c *gin.Context) {
	var queryStruct problemQuery
	err := c.ShouldBind(&queryStruct)
	if err != nil {
		model.ServerLog(err.Error())
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
				"errCode": common.InvalidForm,
				"message": "InvalidForm",
			})
		return
	}
	data, maxPage, err := model.SelectProblemList(queryStruct.Page, queryStruct.Capacity)
	if err != nil {
		model.ServerLog(err.Error())
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
				"errCode": common.DataBaseUnavaliable,
				"message": "Database unavaliable",
			})
		return
	}
	ret, err := json.Marshal(problemListRetWrapper(data, maxPage))
	if err != nil {
		model.ServerLog(err.Error())
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
				"errCode": common.JSONComponentUnavaliable,
				"message": "Json Marshal error",
			})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"errCode": common.Success,
		"message": ret,
	})
}
