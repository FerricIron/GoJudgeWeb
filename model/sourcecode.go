package model

type SourceCode struct {
	Scid  int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Source string `gorm:"type:varchar(4096);NOT NULL"`
	Language int `gorm:"type:int; NOT NULL"`
}

