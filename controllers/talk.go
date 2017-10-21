package controllers

import (
	"fmt"

	"github.com/astaxie/beego"

	"net/http"

	"github.com/lifeisgo/shrimptalk/models"
	"github.com/satori/go.uuid"
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

func (c *TalkController) FindAll() {
	fmt.Println("findallgaoqiankun")

	defer c.ServeJSON()
	c.Data["json"] = map[string]interface{}{
		"message": "ok",
		"data": map[string]interface{}{
			"data": "hao",
		},
	}
}
func (c *TalkController) New() {
	s, _ := models.Session().SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer s.SessionRelease(c.Ctx.ResponseWriter)
	id, b := s.Get("login").(string)
	if !b {
		c.Ctx.Redirect(http.StatusFound, "/")
		return
	}
	user := models.FindUser(id)
	if user.IsNil() {
		c.Ctx.Redirect(http.StatusFound, "/")
		return
	}

	c.Data["NickName"] = user.NickNameHex
}

func (c *TalkController) PostNew() {
	s, _ := models.Session().SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer s.SessionRelease(c.Ctx.ResponseWriter)
	id, b := s.Get("login").(string)
	if !b {
		c.Ctx.Redirect(http.StatusFound, "/")
		return
	}
	user := models.FindUser(id)
	if user.IsNil() {
		c.Ctx.Redirect(http.StatusFound, "/")
		return
	}

	detail := c.GetString("detail")

	talk := models.NewTalk()
	talk.UserID = uuid.FromStringOrNil(id)
	//talk.TalkNameHex = user.NickNameHex
	talk.Now = models.RandomUser().ID
	talk.AddComment(user.NickNameHex, detail)
	talk.Create()
	c.Ctx.WriteString("")
	return

}
