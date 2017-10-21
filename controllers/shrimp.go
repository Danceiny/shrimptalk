package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//	c.Data["Website"] = "beego.me"
	//	c.Data["Email"] = "astaxie@gmail.com"
	//	c.TplName = "index.tpl"
	//	fmt.Println("gaoqiankun")
	fmt.Println("Test")
	defer c.ServeJSON()
	c.Data["json"] = map[string]interface{}{
		"ok": true,
		"data": map[string]interface{}{
			"data":     "恭喜你注册成功",
			"username": "用户名为:",
			"password": "账号密码为:"}}

}
