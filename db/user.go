package db
func login(username,password string)(bool bool,err error){
	ret,err:=db.Query(`SELECT password FROM USER WHERE username=?;`,username)
	if err!=nil{
		return false,err
	}
	col,err:=ret.Columns()
	if err!=nil {
		return false,err
	}
	if col[0]!=password{
		return false,nil
	}
	return true,nil
}
func register(username,password,comment string, schoolID,privilege int,)(bool bool,err error){
	ret,err:=db.Exec("INSERT INTO user(username,password,schoolID,comment,privilege) VALUES(?,?,?,?,?)",username,password,schoolID,comment,privilege)
	if err!=nil{
		return false,err
	}
	affect,err:=ret.RowsAffected()
	if err!=nil||affect==0{
		return false,err
	}
	return true,nil
}
