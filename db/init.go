package db

import (
	"fmt"
    _"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)
const username 	string=	"judgeweb"
const password  string= "judgeweb"
const dbname	string= "judgeweb"
const dbadress	string=	"db.lengyu.me"
const dbport	string=	"3306"
var DSN string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",username,password,dbadress,dbport,dbname)

func init() {
	db,err:=gorm.Open("mysql",DSN)
	if err!=nil{
		panic("Can not connect database.Check config plz\n")
	}
	//check table exit
	if !db.HasTable(&User{}){
		db.CreateTable(&User{})
	}
	if !db.HasTable(&School{}){
		db.CreateTable(&School{})
	}
}
func openConnect()(db *gorm.DB,err error){
	db,err=gorm.Open("mysql",DSN)
	if err!=nil{
		panic("TEST DATABASE CONNECT ERROR!\n")
	}
	return
}
