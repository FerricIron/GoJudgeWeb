package db

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)
type User struct{
	Uid 	 	int 	`gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"`
	Username 	string 	`gorm:"type:varchar(20);unique_index;NOT NULL"`
	Nickname 	string 	`gorm:"type:varchar(20);NOT NULL"`
	Password 	string 	`gorm:"type:char(32);NOT NULL"`
	Description	string 	`gorm:"type:varchar(255)"`
	School		School	`gorm:"ForeignKey:sid;AssociationForeignKey:sid"`
	Sid 		int
	Privilege	int 	`gorm:"type:int;NOT NULL"`
	Submitcount	int		`gorm:"type:int;NOT NULL"`
	Solved		int 	`gorm:"type:int;NOT NULL"`
}
func pass2md5(password string)(md5Password string){
	md5Ctx:=md5.New()
	md5Ctx.Write([]byte(password))
	md5Password=fmt.Sprint(hex.EncodeToString(md5Ctx.Sum(nil)))
	return
}
func Test(){
	user:=User{
		Username:"lengyu",
		Nickname:"test",
		Password:pass2md5("test"),
		Privilege:0,
		Submitcount:0,
		Solved:0,
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