package controllers

import (
	"github.com/astaxie/beego"
	"github.com/lifeisgo/shrimptalk/models"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	id := c.Ctx.Input.Param(":id")
	s, _ := models.Session().SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer s.SessionRelease(c.Ctx.ResponseWriter)
	user := models.User{}
	models.ORM().Where("nick_name_hex = ?", id).Find(&user)
	if user.IsNil() {
		c.Ctx.WriteString("用户不存在!")
		return
	}
	if l, b := s.Get("login").(string); b && l == user.ID.String() {
		c.Ctx.WriteString(id + " 已经登录!")
		return
	}

	s.Set("login", user.ID.String())
	c.Ctx.WriteString("登陆中 " + id + " !")
}
