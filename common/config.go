package common

import (
	"github.com/go-ini/ini"
	"os"
	"syscall"
	"time"
)

const configPath  = "./config.ini"
const watchingSleepTime = 5 //second
type JudgeServer struct {
	Address string	`ini:"address"`
	Port   string	`ini:"port"`
}
type WebServer struct {
	Address string	`ini:"address"`
	Port   string	`ini:"port"`
}
type DataBase struct {
	Address   string	`ini:"address"`
	Port     string		`ini:"port"`
	Name     string		`ini:"name"`
	Username string		`ini:"username"`
	Password string		`ini:"password"`
}

type Config struct {
	JudgeServer
	WebServer
	DataBase
}
var GlobalConfig *Config = new(Config)
func init(){

}
func ParseConfig() {
	cfg,err:=ini.Load(configPath)
	if err!=nil{
		panic(err.Error())
	}
	if err=cfg.MapTo(GlobalConfig);err!=nil{
		panic(err)
	}
}
func ConfigFileWatching(){
	var lastChangeTime int64 = 0
	for ;;{
		fileInfo,err:=os.Stat(configPath)
		if err!=nil{
			time.Sleep(watchingSleepTime*time.Second)
			continue
		}
		stat_t:=fileInfo.Sys().(*syscall.Stat_t)
		changeTime :=syscall.TimespecToNsec(stat_t.Mtim)
		if lastChangeTime==0{
			lastChangeTime=changeTime
		}else {
			if changeTime>lastChangeTime {
				lastChangeTime=changeTime
				ParseConfig()
				// there need log , but I want to moyu
			}
			time.Sleep(watchingSleepTime*time.Second)
		}
	}
}