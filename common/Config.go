package common

import (
	"github.com/go-ini/ini"
)

const configPath  = "./config.ini"
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
