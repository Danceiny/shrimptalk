package controllers

import (
	"github.com/astaxie/beego"

	"github.com/lifeisgo/shrimptalk/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) URLMapping() {
	c.Mapping("All", c.All)
}

func (c *MainController) Get() {
	c.Ctx.WriteString("hello world")
}

func (c *MainController) All() {
	tlist := []models.Talk{}
	models.ORM().Table("talks").Find(&tlist).Order("by max desc")
	//for _, v := range tlist {
	//	c.Ctx.WriteString(v.ToString())
	//}
	c.Data["Talk"] = tlist
	c.TplName = "talk_all.tpl"
}
