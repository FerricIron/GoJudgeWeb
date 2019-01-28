package model

type Problem struct {
	ProblemId   int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	ProblemName string `gorm:"type:varchar(128);NOT NULL"`
	Author      string `gorm:"type:varchar(32);NOT NULL"`
	Description string `gorm:"type:varchar(256);NOT NULL"`
	Property    string `gorm:"type:int;NOT NULL"`
	Privilege	int    `gorm:"type:int;NOT NULL"`
	SubmitCount string `gorm:"type:int;NOT NULL"`
	Solved      string `gorm:"type:int;NOT NULL"`
	TimeLimit   string `gorm:"type:int;NOT NULL"`
	MemoryLimit string `gorm:"type:int;NOT NULL"`
	Status      string `gorm:"type:int;NOT NULL"`
}

func InsertProblem(problem *Problem) error {
	db, err := openConnect()
	defer db.Close()
	if err != nil {
		ServerLog(err.Error())
		return err
	}
	return db.Create(problem).First(problem).Error
}
func SelectProblemList(page,capacity int)(data []Problem,maxPage int,err error){
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
	err=db.Table("problems").Count(&maxPage).Error
	maxPage=maxPage/capacity
	if err!=nil{
		return nil,0,err
	}
	err=db.Offset(offset).Find(&data).Limit(capacity).Error
	if err!=nil{
		ServerLog(err.Error())
		return nil,0,err
	}
	return
}
func SelectProblem(ProblemID int) (problem Problem, err error) {
	db, err := openConnect()
	defer db.Close()
	if err != nil {
		ServerLog(err.Error())
		return Problem{}, nil
	}
	err=db.Where("problem_id = ? ",ProblemID).First(&problem).Error
	return
}

