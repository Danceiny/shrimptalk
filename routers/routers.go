package routers

import (
	"github.com/astaxie/beego"
	"github.com/lifeisgo/shrimptalk/controllers"
)

func init() {
	beego.Router("/all", &controllers.MainController{}, "get:All")
}
