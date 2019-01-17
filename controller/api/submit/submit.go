package submit

import (
	"github.com/ferriciron/GoJudgeWeb/common"
	"github.com/ferriciron/GoJudgeWeb/controller/api"
	"github.com/ferriciron/GoJudgeWeb/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type submitForm struct {
	Language   	int 		`form:"language" binding:"required"`
	SourceCode 	string 		`form:"source" binding:"required"`
	ProblemID  	int 		`form:"problemid" binding:"required"`
	ContestID	int			`form:"contestid"`
}
func submitToJudgeServer(){
	return
}
func Submit(c *gin.Context) {
	var submit submitForm
	if err := c.ShouldBind(&submit); err != nil {
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
				"errCode": common.InvalidForm,
				"message": err.Error(),
			})
		return
	}
	token := c.GetHeader("token")
	var jwt common.JWT
	claims, err := jwt.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
				"errCode": common.TokenInvalid,
				"message": err.Error(),
			})
		return
	}
	problem, err := model.SelectProblem(submit.ProblemID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
				"errCode": common.DataBaseUnavaliable,
				"message": err.Error(),
			})
		return
	}
	if !api.CheckPrivilege(claims.Privilege,problem.Privilege){
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
			"errCode":common.PermissionDenied,
			"message":"Permission denied",
		})
		return
	}
	sourceCode:=model.SourceCode{
		Source:   submit.SourceCode,
		Language: submit.Language,
	}
	err=model.InsertSourceCode(&sourceCode)
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
				"errCode":common.DataBaseUnavaliable,
				"message":err.Error(),
			})
		return
	}
	submitModel :=model.Submit{
		Uid:        claims.UID,
		Time:       int(time.Now().Unix()),
		Language:   submit.Language,
		SourceCode: model.SourceCode{},
		Scid:       sourceCode.Scid,
		ContestId:  submit.ContestID,
		ProblemId:  submit.ProblemID,
		Status:     0,
	}
	err=model.InsertSubmit(&submitModel)
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
				"errCode":common.DataBaseUnavaliable,
				"message":err.Error(),
			})
		return
	}
	c.JSON(http.StatusOK,
		gin.H{
			"errCode":common.Success,
			"message":"ok",
		})
	submitToJudgeServer()
}
