package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "myblog/routers"
	"myblog/utils"
)

func main() {
	utils.InitMysql()
	beego.Run()
}
