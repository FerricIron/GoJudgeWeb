package model

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
