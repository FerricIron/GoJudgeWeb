package email

import (
	"github.com/ferriciron/GoJudgeWeb/common"
	"gopkg.in/gomail.v2"
)

func StartListen() *gomail.Dialer {
	if common.GlobalConfig.EmailServer.Status==0{
		return nil
	}
	daemon := gomail.NewDialer(common.GlobalConfig.EmailServer.SMTPAdress, common.GlobalConfig.EmailServer.SMTPPort, common.GlobalConfig.EmailServer.Username, common.GlobalConfig.EmailServer.Password)
	return daemon
}
func init(){

}
