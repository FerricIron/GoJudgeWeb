package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
const username 	string=	""
const password  string= 	""
const dbname	string= 	""
var db *sql.DB
func init(){
	temp:=username+":"+password+"@/"+dbname
	var err error
	db,err=sql.Open("mysql",temp)
	if err!=nil {
		panic("open database error")
	}
}
