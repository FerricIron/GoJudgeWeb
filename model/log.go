package model

import (
	"encoding/json"
	"fmt"
	"github.com/ferriciron/GoJudgeWeb/common"
	"github.com/gin-gonic/gin"
	"log"
	"path/filepath"
	"runtime"
	"time"
)

type Log struct {
	Time       time.Time
	UserIP     string `gorm:"size:18"`
	HttpMethod string
	URI        string
	Header     string
	Message    string
	TracePath  string
	Level      int
}

func insertLog(logInfo Log) {
	db, err := openConnect()
	defer db.Close()
	if err != nil {
		log.Print("InsertLog:" + err.Error() + "\n")
	}
	err = db.Create(&logInfo).Error
	if err != nil {
		log.Print("InsertLog:" + err.Error() + "\n")
	}
}

func baseLog(c *gin.Context, message string, level int) {
	filename, line, funcname := "???", 0, "???"
	pc, filename, line, ok := runtime.Caller(1)
	if ok {
		funcname = runtime.FuncForPC(pc).Name()
		filename = filepath.Base(filename)
	}
	logModel := Log{
		Time:       time.Now(),
		UserIP:     c.ClientIP(),
		HttpMethod: c.Request.Method,
		URI:        c.Request.RequestURI,
		Message:    message,
		TracePath:  fmt.Sprintf("%s:%d:%s\n", filename, line, funcname),
		Level:      level,
	}
	data, err := json.Marshal(c.Request.Header)
	if err != nil {
		logModel.Header = ""
	} else {
		logModel.Header = string(data)
	}
	if common.GlobalConfig.WebServer.Debug == 1 {
		fmt.Println(logModel)
	} else {
		insertLog(logModel)
	}
}
func InfoLog(c *gin.Context, message string) {
	baseLog(c, message, 3)
}
func WarnLog(c *gin.Context, message string) {
	baseLog(c, message, 2)
}
func ErrLog(c *gin.Context, message string) {
	baseLog(c, message, 1)
}


type ServerRuntimeLog struct {
	Time time.Time
	Message    string
	TracePath  string
}
func insertServerRuntimeLog(runtimeLog ServerRuntimeLog){
	db,err:=openConnect()
	defer db.Close()
	if err!=nil{
		log.Printf("InsertServerRuntime:%s\n",err)
	}
	err=db.Create(&runtimeLog).Error
	if err!=nil{
		log.Printf("InsertServerRuntime:%s\n",err)
	}
}
func ServerLog(message string){
	filename, line, funcname := "???", 0, "???"
	pc, filename, line, ok := runtime.Caller(1)
	if ok {
		funcname = runtime.FuncForPC(pc).Name()
		filename = filepath.Base(filename)
	}
	logModel := ServerRuntimeLog{
		Time:      time.Now(),
		Message:   message,
		TracePath: fmt.Sprintf("%s:%d:%s\n", filename, line, funcname),
	}
	if common.GlobalConfig.WebServer.Debug==1{
		log.Println(logModel)
	}else {
		insertServerRuntimeLog(logModel)
	}
}