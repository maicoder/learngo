package model

import (
	"imooc.com/ccmouse/learngo/learngin/ginmysql/initDB"
	"log"
)

type UserModel struct {
	Email         string `form:"email"`
	Password      string `form:"password"`
	PasswordAgain string `form:"password-again"`
}

func (user *UserModel) Save() int64 {
	result, err := initDB.Db.Exec("insert into ")
	if err != nil {
		log.Panicln("user insert error", err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Panicln("user insert id user", err.Error())
	}
	return id
}