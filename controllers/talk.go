package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

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
