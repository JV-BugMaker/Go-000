package main

import "fmt"

type ErrNotFound struc{
	code int
	msg string
	err error
}


// dao层出现的error
func dao(model *db.Model)error{
	user := new(User)
	err := DB.QueryRow("select id,username,password from users where id=?", id).Scan(user)
	if err != nil {
		return errors.Wrapf(ErrNotFound, fmt.Sprintf("sql: %s error: %v", sql, err))
	}
}


func biz(){
	model,err := dao(&user)
	if errors.Is(err,ErrNotFound){
		// 处理
	}
}
