package model

type SourceCode struct {
	Scid     int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Source   string `gorm:"type:varchar(4096);NOT NULL"`
	Language int    `gorm:"type:int; NOT NULL"`
}
func InsertSourceCode(code *SourceCode)(err error){
	db,err:=openConnect()
	defer db.Close()
	if err!=nil{
		return
	}
	return db.Create(code).First(code).Error
}
