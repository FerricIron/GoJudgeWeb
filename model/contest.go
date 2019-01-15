package model

type Contest struct {
	ContestId   int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	ContestName string `gorm:"type:varchar(56);NOT NULL"`
	Creator     string `gorm:"type:varchar(128);NOT NULL"`
	Desciption  string `gorm:"type:varchar(256);"`
	StartTime   int    `gorm:"type:int;NOT NULL"`
	EndTime     int    `gorm:"type:int;NOT NULL"`
	Property    int    `gorm:"type:int;NOT NULL"`
	Privilege   int    `gorm:"type:int;NOT NULL"`
	Status      int    `gorm:"type:int;NOT NULL"`
}
type ContestInfo struct {
	Contest     Contest `gorm:"foreignkey:ContestId"`
	ContestId   int     `gorm:"NOT NULL"`
	Problem     Problem `gorm:"foreignkey:ProblemId"`
	ProblemId   int     `gorm:"NOT NULL"`
	Id          int     `gorm:"type:int;NOT NULL"`
	SubmitCount int     `gorm:"type:int;NOT NULL"`
	Solved      int     `gorm:"type:int;NOT NULL"`
	Penalty     int     `gorm:"type:int;NOT NULL"`
}

type ContestRegister struct {
	Contest     Contest `gorm:"foreignkey:ContestId"`
	ContestId   int     `gorm:"NOT NULL"`
	User        User    `gorm:"foreignkey:Uid"`
	Uid         int     `gorm:"NOT NULL"`
	SubmitCount int     `gorm:"type:int;NOT NULL"`
	Solved      int     `gorm:"type:int;NOT NULL"`
	Penalty     int     `gorm:"type:int;NOT NULL"`
	Rank        int     `gorm:"type:int;NOT NULL"`
}
