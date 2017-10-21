package main

import (
	"github.com/astaxie/beego"
	_ "github.com/lifeisgo/shrimptalk/routers"
)

func main() {
	beego.Run()
}
