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

func (c *MainController) All() {
	tlist := []models.Talk{}
	models.ORM().Table("talks").Order("max desc").Find(&tlist)
	c.Data["Talk"] = tlist

}
