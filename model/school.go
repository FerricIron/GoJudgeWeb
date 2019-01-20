package model

type School struct {
	Sid       int    `gorm:"type:int;PRIMARY_KEY;AUTO_INCREMENT"`
	Name      string `gorm:"type:varchar(32);NOT NULL;unique;index"`
	ShortName string `gorm:"type:varchar(32);NOT NULL"`
}
