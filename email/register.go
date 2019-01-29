package email

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/ferriciron/GoJudgeWeb/common"
	"github.com/ferriciron/GoJudgeWeb/model"
	"gopkg.in/gomail.v2"
	"time"
)
const salt = "GfsatVfisgk"
func GenerateContent(username,password,email string,t int64)string{
	md5Ctx:=md5.New()
	md5Ctx.Write([]byte(fmt.Sprintf("%s%s%s%s%d",username,salt,password,email,t)))
	return fmt.Sprint(hex.EncodeToString(md5Ctx.Sum(nil)))
}
func SendRegisterEmail(User model.User,email string){
	uid:=User.Uid
	fmt.Print(uid)
	username:=User.Username
	password:=User.Password
	t:=int64(time.Now().Unix())
	md5Content:=GenerateContent(username,password,email,t)
	registerConfirmUrl:=fmt.Sprintf("%s/data/user/%d/email/confirm?t=%d&c=%s&e=%s",common.GlobalConfig.EmailServer.Export,uid,t,md5Content,email)
	m := gomail.NewMessage()
	m.SetHeader("From", "email@lengyu.me")
	m.SetHeader("To", "727109404@qq.com")
	m.SetHeader("Subject", "Register confirm")
	m.SetBody("text/html", registerConfirmUrl)
	go sendEmail(m)
}
func TestSend(){
	User:= model.User{
		Uid:         1,
		Username:    "Maymomoo",
		Nickname:    "test",
		Password:    "e11170b8cbd2d74102651cb967fa28e5",
		Description: "test",
		Sid:         0,
		Privilege:   0,
		SubmitCount: 0,
		Solved:      0,
	}
	SendRegisterEmail(User,"727109404@qq.com")
}
