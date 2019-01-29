package model

type School struct {
	Sid       int    `gorm:"type:int;PRIMARY_KEY;AUTO_INCREMENT"`
	Name      string `gorm:"type:varchar(64);NOT NULL;unique;index"`
	ShortName string `gorm:"type:varchar(32);NOT NULL"`
}

func SelectAllSchool()(data []School,err error){
	db,err:=openConnect()
	defer db.Close()
	if err!=nil{
		ServerLog(err.Error())
		return nil,err
	}
	err=db.Find(&data).Error
	return
}