package controllers

import "github.com/astaxie/beego"

type TalkController struct {
	beego.Controller
}

func (c *TalkController) URLMapping() {
	c.Mapping("Talk", c.Talk)
}

func (c *TalkController) Talk() {
	c.Ctx.WriteString(c.Ctx.Input.Param(":id"))
	c.Ctx.WriteString("talk:")
	c.Ctx.WriteString(c.Ctx.Input.Param(":talk"))

}
