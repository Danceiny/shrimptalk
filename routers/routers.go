package routers

import (
	"github.com/lifeisgo/shrimptalk/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.UserController{}, "post:Register")
	beego.Router("/all", &controllers.MainController{}, "get:All")
	beego.Router("/:id/talk/:talk", &controllers.TalkController{}, "get:Talk")
	beego.Router("/new", &controllers.TalkController{}, "get:New")
	beego.Router("/:id/new", &controllers.TalkController{}, "post:PostNew")
	beego.Router("/login/:id", &controllers.LoginController{})
}
