package model

import "math"

type Submit struct {
	SubmitId   int        `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	User       User       `gorm:"foreignkey:Uid;"`
	Uid        int        `gorm:"type:int;NOT NULL"`
	Time       int        `gorm:"type:int;NOT NULL"`
	Language   int        `gorm:"type:int;NOT NULL"`
	SourceCode SourceCode `gorm:"foreignkey:Scid"`
	Scid       int        `gorm:"type:int;NOT NULL"`
	Contest    Contest    `gorm:"foreginkey:ContestId;"`
	ContestId  int        `gorm:"type:int;NOT NULL"`
	Problem    Problem    `gorm:"foreginkey:ProblemId"`
	ProblemId  int        `gorm:"type:int;NOT NULL"`
	Status     int        `gorm:"type:int;NOT NULL"`
	TimeCost   int        `gorm:"type:int"`
	Info       string     `gorm:"type:varchar(256)"`
}
func InsertSubmit(submit *Submit)(err error){
	db,err:=openConnect()
	defer db.Close()
	if err!=nil{
		return
	}
	return db.Create(submit).First(submit).Error
}
func SelectStatusList(page,capacity int)(data []Submit,maxPage int,err error) {
	if capacity<1 {
		capacity=1
	}else if capacity>100{
		capacity=100
	}

	db,err:=openConnect()
	defer db.Close()
	if err!=nil{
		ServerLog(err.Error())
		return nil,0,err
	}
	offset:=(page-1)*capacity
	err=db.Table("submits").Count(&maxPage).Error
	maxPage=int(math.Ceil(float64(maxPage)/float64(capacity)))
	if err!=nil{
		return nil,0,err
	}
	err=db.Order("submit_id desc").Offset(offset).Find(&data).Limit(capacity).Error
	if err!=nil{
		ServerLog(err.Error())
		return nil,0,err
	}
	return
}