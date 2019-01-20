package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ferriciron/GoJudgeWeb/common"
)

func init() {
	common.ParseConfig()
	initTables()
}
func getDSN()string{
	config:=common.GlobalConfig
	var DSN string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		config.DataBase.Username,config.DataBase.Password,config.DataBase.Address,config.DataBase.Port,config.DataBase.Name)
	return DSN
}

func initTables() {
	db, err := gorm.Open("mysql", getDSN())
	defer db.Close()
	if err != nil {
		panic("Can not connect database.Check config plz\n")
	}
	db.AutoMigrate(&School{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Problem{})
	db.AutoMigrate(&Submit{})
	db.AutoMigrate(&Contest{})
	db.AutoMigrate(&ContestInfo{})
	db.AutoMigrate(&ContestRegister{})
	db.AutoMigrate(&SourceCode{})

}
func openConnect() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", getDSN())
	if err != nil {
		panic("TEST DATABASE CONNECT ERROR!\n")
	}
	return
}
