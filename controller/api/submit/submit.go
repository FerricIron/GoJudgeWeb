package submit

import (
	"encoding/json"
	"fmt"
	"github.com/ferriciron/GoJudgeWeb/common"
	"github.com/ferriciron/GoJudgeWeb/controller/api"
	"github.com/ferriciron/GoJudgeWeb/model"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"time"
)

type submitForm struct {
	Language   int    `form:"language" binding:"required"`
	SourceCode string `form:"source" binding:"required"`
	ProblemID  int    `form:"problemid" binding:"required"`
	ContestID  int    `form:"contestid"`
}
type submitStruct struct {
	SubmitID   int    `json:"submitID"`
	ProblemID  int    `json:"problemId"`
	CodeSource []byte `json:"codeSource"`
	Language   int    `json:"language"`
}
type responseStruct struct {
	ErrCode   int    `json:"errCode"`
	JudgeNode int    `json:"judgeNode"`
	AllNode   int    `json:"allNode"`
	TimeCost  int    `json:"timecost"`
	Msg       []byte `json:"msg"`
}

func (resp *responseStruct) StructToBytes() (data []byte, err error) {
	data, err = json.Marshal(resp)
	return
}

func (submit *submitStruct) StructToBytes() (data []byte, err error) {
	data, err = json.Marshal(submit)
	return
}

func submitToJudgeServer(submit submitStruct) {
	listenAddress := fmt.Sprintf("%s:%s", common.GlobalConfig.JudgeServer.Address, common.GlobalConfig.JudgeServer.Port)
	net, err := net.Dial("tcp", listenAddress)
	defer net.Close()
	// need log
	if err != nil {
		fmt.Print(err.Error())
	}
	socket := common.SocketFromConn(net)
	encoder := common.NewEnCoder()
	decoder := common.NewDecoder()
	err = encoder.SendStruct(socket, &submit)
	// need log
	if err != nil {
		fmt.Print(err.Error())
	}
	var response responseStruct
	for {
		err = decoder.ReadStruct(socket, &response)
		//need logs
		if err != nil {
			fmt.Print(err)
			break
		} else {
			if response.JudgeNode == response.AllNode || response.ErrCode != common.AcceptCode {

				return
			}
			//need update
			fmt.Print(response)
		}
	}

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
	if !api.CheckPrivilege(claims.Privilege, problem.Privilege) {
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
				"errCode": common.PermissionDenied,
				"message": "Permission denied",
			})
		return
	}
	sourceCode := model.SourceCode{
		Source:   submit.SourceCode,
		Language: submit.Language,
	}
	err = model.InsertSourceCode(&sourceCode)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
				"errCode": common.DataBaseUnavaliable,
				"message": err.Error(),
			})
		return
	}
	submitModel := model.Submit{
		Uid:        claims.UID,
		Time:       int(time.Now().Unix()),
		Language:   submit.Language,
		SourceCode: model.SourceCode{},
		Scid:       sourceCode.Scid,
		ContestId:  submit.ContestID,
		ProblemId:  submit.ProblemID,
		Status:     0,
	}
	err = model.InsertSubmit(&submitModel)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK,
			gin.H{
				"errCode": common.DataBaseUnavaliable,
				"message": err.Error(),
			})
		return
	}
	c.JSON(http.StatusOK,
		gin.H{
			"errCode": common.Success,
			"message": "ok",
		})
	var postData submitStruct
	postData.ProblemID = submitModel.ProblemId
	postData.CodeSource = []byte(submit.SourceCode)
	postData.Language = submit.Language
	postData.SubmitID = submitModel.SubmitId
	go submitToJudgeServer(postData)
}
