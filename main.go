package main

import (
	"github.com/astaxie/beego"
	//	"github.com/astaxie/beego/logs"

	_ "github.com/lifeisgo/shrimptalk/routers"
)

func main() {
	//	logs.SetLogger("console")
	beego.Run()
}
