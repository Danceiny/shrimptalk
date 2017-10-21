package routers

import (
	"github.com/lifeisgo/shrimptalk/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/all", &controllers.MainController{}, "get:All")

}
