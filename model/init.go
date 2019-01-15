package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const username string = "judgeweb"
const password string = "judgeweb"
const dbname string = "judgeweb"
const dbadress string = "db.lengyu.me"
const dbport string = "3306"

var DSN string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", username, password, dbadress, dbport, dbname)

func init() {
	db, err := gorm.Open("mysql", DSN)
	defer db.Close()
	if err != nil {
		panic("Can not connect database.Check config plz\n")
	}
	//check table exit
	db.AutoMigrate(&School{})
	db.CreateTable(&User{})

}
func openConnect() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", DSN)
	if err != nil {
		panic("TEST DATABASE CONNECT ERROR!\n")
	}
	return
}
