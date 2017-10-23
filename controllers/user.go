package controllers

import (

	//	"net/http"

	"github.com/astaxie/beego"
	"github.com/lifeisgo/shrimptalk/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Register() {
	user := models.NewUser()
	models.AddUser(user)
	message := "恭喜  " + user.NickNameHex + " 注册成功！！！！"
	c.Ctx.WriteString(message)
}
