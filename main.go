package main

import (
	"github.com/astaxie/beego"
	_ "github.com/lifeisgo/shrimptalk/routers"

	//	"fmt"

	_ "github.com/lifeisgo/shrimptalk/common"
)

func main() {
	beego.Run()

}
