package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type User struct {
	Uid         int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"`
	Username    string `gorm:"type:varchar(20);unique_index;NOT NULL"`
	Nickname    string `gorm:"type:varchar(20);NOT NULL"`
	Password    string `gorm:"type:char(32);NOT NULL"`
	Description string `gorm:"type:varchar(255)"`
	School      School `gorm:"ForeignKey:sid;AssociationForeignKey:sid"`
	Sid         int
	Privilege   int `gorm:"type:int;NOT NULL"`
	Submitcount int `gorm:"type:int;NOT NULL"`
	Solved      int `gorm:"type:int;NOT NULL"`
}

func pass2md5(password string) (md5Password string) {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	md5Password = fmt.Sprint(hex.EncodeToString(md5Ctx.Sum(nil)))
	return
}

func Test() {
	user := User{
		Username:    "lengyu",
		Nickname:    "test",
		Password:    pass2md5("test"),
		Privilege:   0,
		Submitcount: 0,
		Solved:      0,
	}
	db, err := openConnect()
	if err != nil {
		fmt.Print("db error")
	}
	db.NewRecord(user)
	db.Create(&user)
}
