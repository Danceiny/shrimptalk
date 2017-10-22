package controllers

import (
	"fmt"

	"github.com/astaxie/beego"

	"net/http"

	"log"

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

func (c *TalkController) FindMyTalk() {
	fmt.Println("高乾坤")
	s, _ := models.Session().SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer s.SessionRelease(c.Ctx.ResponseWriter)
	id, b := s.Get("login").(string)
	if !b {
		c.Ctx.Redirect(http.StatusFound, "/")
		return
	}

	talks := models.FindAllTalk(id)

	c.Data["Talks"] = talks

}

func (c *TalkController) FindNowTalk() {

	fmt.Println("我收到的")
	s, _ := models.Session().SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer s.SessionRelease(c.Ctx.ResponseWriter)
	id, b := s.Get("login").(string)
	if !b {
		c.Ctx.Redirect(http.StatusFound, "/")
		return
	}
	//	id := "b5712894-b640-11e7-a2bd-acbc32a50041"
	//	user := models.FindUser(id)
	mytalks := models.FindByNow(id)
	//	mytalks := models.FindByNow("8ea685af-b649-11e7-ac6e-acbc32a50041")
	c.Data["Talks"] = mytalks

}
func (c *TalkController) Answer() {
	fmt.Println("answer!!!!")
	s, _ := models.Session().SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer s.SessionRelease(c.Ctx.ResponseWriter)
	id, b := s.Get("login").(string)
	if !b {
		c.Ctx.Redirect(http.StatusFound, "/")
		return
	}
	talkNameHex := c.Ctx.Input.Param(":id")
	fmt.Println("id:", talkNameHex)

	talk := models.FindByTalkNameHex(talkNameHex)
	fmt.Println("talk:", talk.ToComment())
	user := models.FindUser(id)
	if user.IsNil() {
		c.Ctx.Redirect(http.StatusFound, "/")
		return
	}
	c.Data["NickName"] = user.NickNameHex
	c.Data["talk"] = talk.ToComment()
	//	c.Data["user"] = user
}
func (c *TalkController) PostAnswer() {
	fmt.Println("回答ßßß!!!!")
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
	tk := new(models.Talk)
	tk2 := new(models.Talk)
	tk.TalkNameHex = c.Ctx.Input.Param(":talkhex")
	models.ORM().Where(tk).FirstOrInit(tk2)
	tk2.AddComment(user.NickNameHex, detail)
	next := models.RandomUser()
	tk2.Now = next.ID
	models.ORM().Where(tk).Save(tk2)
	c.Data["Next"] = next
	//	c.Ctx.Redirect(http.StatusFound, "/talk/mytalk")
	//	fmt.Println("detail:", detail, id)
	//	c.Ctx.WriteString(detail)
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
	next := models.RandomUser()
	talk.Now = next.ID
	talk.AddComment(user.NickNameHex, detail)
	talk.Create()
	c.Data["Detail"] = detail
	c.Data["Next"] = next
	return

}

func (c *TalkController) Detail() {
	id := c.Ctx.Input.Param(":id")
	talk := models.FindTalkByHex(id)
	c.Data["Detail"] = talk.ToComment()
	log.Println(talk)
	user := models.FindUser(talk.Now.String())
	c.Data["Next"] = user
}
