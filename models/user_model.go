package models

import (
	"fmt"
	"myblog/utils"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int //0 正常状态 1 删除
	Createtime int64
}

func InsertUser(user User) (int64, error) {
	return utils.ModifyDB("insert into users(username,password,status,createtime) values (?,?,?,?)", user.Username, user.Password, user.Status, user.Createtime)
}

func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username = '%s'", username)
	return QueryUserWithCon(sql)
}

func QueryUserWithCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}
