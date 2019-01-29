package email

import (
	"github.com/ferriciron/GoJudgeWeb/common"
	"github.com/ferriciron/GoJudgeWeb/model"
	"gopkg.in/gomail.v2"
)

func sendEmail(message *gomail.Message){
	if common.GlobalConfig.EmailServer.Status==0{
		return
	}
	d:=StartListen()
	if err:=d.DialAndSend(message);err!=nil{
		model.ServerLog(err.Error())
	}
}

