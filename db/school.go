package db
type School struct{
	sid			int 	`gorm:"type:int;PRIMARY_KEY;AUTO_INCREMENT"`
	name		string 	`gorm:"type:varchar(32);NOT NULL;unique;index"`
	shortname	string 	`gorm:"type:varchar(32);NOT NULL"`
}
