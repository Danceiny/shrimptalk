package controllers

import (
	"net/http"

	"log"

	"github.com/astaxie/beego"
	"github.com/lifeisgo/shrimptalk/models"
)

type TalkController struct {
	beego.Controller
}

func (c *TalkController) URLMapping() {
	c.Mapping("Talk", c.Talk)
	c.Mapping("New", c.New)
	c.Mapping("PostNew", c.PostNew)
}

func (c *TalkController) Talk() {
	c.Ctx.WriteString(c.Ctx.Input.Param(":id"))
	c.Ctx.WriteString("talk:")
	c.Ctx.WriteString(c.Ctx.Input.Param(":talk"))

}

func (c *TalkController) New() {
	s, _ := models.Session().SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer s.SessionRelease(c.Ctx.ResponseWriter)
	id := s.Get("login")
	user := models.FindUser(id.(string))
	if user.IsNil() {
		c.Ctx.Redirect(http.StatusOK, "/")
	}

	c.Data["NickName"] = user.NickNameHex
}

func (c *TalkController) PostNew() {
	s, _ := models.Session().SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer s.SessionRelease(c.Ctx.ResponseWriter)
	id := s.Get("login")
	user := models.FindUser(id.(string))
	if user.IsNil() {
		c.Ctx.Redirect(http.StatusOK, "/")
	}

	c.Data["NickName"] = user.NickNameHex
	log.Println("postnew")
	c.Ctx.WriteString("new received")
	c.StopRun()
}
