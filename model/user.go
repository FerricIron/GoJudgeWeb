package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type User struct {
	Uid         int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY" json:"uid"`
	Username    string `gorm:"type:varchar(20);unique_index;NOT NULL" json:"username"`
	Nickname    string `gorm:"type:varchar(20);NOT NULL" json:"nickname"`
	Password    string `gorm:"type:char(32);NOT NULL" json:"password,omitempty"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	School      School `gorm:"ForeignKey:Sid;" json:"school"`
	Email		string `gorm:"type:varchar(64);" json:"email"`
	Sid         int    `gorm:"type:int;NOT NULL" json:"sid"`
	Privilege   int    `gorm:"type:int;NOT NULL" json:"privilege"`
	SubmitCount int    `gorm:"type:int;NOT NULL" json:"submitcount"`
	Solved      int    `gorm:"type:int;NOT NULL" json:"solved"`
}

func pass2md5(password string) (md5Password string) {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	md5Password = fmt.Sprint(hex.EncodeToString(md5Ctx.Sum(nil)))
	return
}
func InsertUser(user *User) (err error) {
	user.Password = pass2md5(user.Password)
	db, err := openConnect()
	defer db.Close()
	if err != nil {
		return
	}
	return db.Create(user).First(user).Error
}
func CheckUserPassword(username, password string) (user *User, err error) {
	pass := pass2md5(password)
	db, err := openConnect()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	err = db.Where("username = ? AND password = ?", username, pass).First(&user).Error
	return
}
func UpdateUser(uid int,user *User)(*User,error){
	db,err:=openConnect()
	if err!=nil{
		return nil,err
	}
	var ret User
	err=db.Find(&ret).Where("uid=?",uid).Error
	if err!=nil{
		return nil,err
	}
	err=db.Model(ret).UpdateColumns(user).First(&ret).Error
	return &ret,err
}
func SelectUser(uid int)(user *User,err error){
	db,err:=openConnect()
	if err!=nil{
		return nil,err
	}
	user = new(User	)
	err=db.Find(user).Where("uid=?",uid).Error
	return
}