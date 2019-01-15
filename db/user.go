package db

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)
type User struct{
	uid 	 	int 	`gorm:"type:int;PRIMARY_KEY;AUTO_INCREMENT"`
	username 	string `gorm:"type:varchar(20);unique_index;NOT NULL"`
	nickname 	string `gorm:"type:varchar(20);NOT NULL"`
	password 	string `gorm:"type:char(32);NOT NULL"`
	description	string `gorm:"type:varchar(255)"`
	school		School	`gorm:"ForeignKey:sid;AssociationForeignKey:sid"`
	sid 		int
	privilege	int 	`gorm:"type:int;NOT NULL"`
	submitcount	int		`gorm:"type:int;NOT NULL"`
	solved		int 	`gorm:"type:int;NOT NULL"`
}
func pass2md5(password string)(md5Password string){
	md5Ctx:=md5.New()
	md5Ctx.Write([]byte(password))
	md5Password=fmt.Sprint(hex.EncodeToString(md5Ctx.Sum(nil)))
	return
}
func Test(){
	user:=User{
		username:"lengyu",
		nickname:"test",
		password:pass2md5("test"),
		privilege:0,
		submitcount:0,
		solved:0,
	}
	db,err:=openConnect()
	if err!=nil {
		fmt.Print("db error")
	}
	db.NewRecord(user)
	db.Create(&user)
}
/*
func login(username,password string)(uid int,err error){
	db,err:=openConnect()
	defer db.Close()
	md5Pass:=pass2md5(password)
	if err!=nil{
		return -1,err
	}
	stmSlc,err:=db.Prepare("SELECT uid FROM user WHERE username=? AND password=?")
	defer stmSlc.Close()
	if err!=nil{
		return -1,err
	}
	err=stmSlc.QueryRow(username,md5Pass).Scan(&uid)
	if err!=nil{
		return -1,err
	}
	return uid,nil
}
func register(username,password string)(bool bool,err error){
	db,err:=openConnect()
	defer db.Close()
	md5Pass:=pass2md5(password)
	if err!=nil{
		return false,err
	}
	stmIns,err:=db.Prepare("INSERT INTO user(username,password) VALUES(?,?)")
	defer stmIns.Close()
	if err!=nil{
		return false,err
	}
	insRet,err:=stmIns.Exec(username,md5Pass)
	if err!=nil{
		return false,err
	}
	i,err:=insRet.RowsAffected()
	if i!=1||err!=nil{
		return false,err
	}
	return true,nil
}*/