package controllers

import (
	"fmt"
	//	"net/http"

	"github.com/astaxie/beego"
	"github.com/lifeisgo/shrimptalk/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
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
func (c *UserController) Register() {
	user := models.NewUser()
	models.AddUser(user)
	message := "恭喜  " + user.NickNameHex + " 注册成功！！！！"
	c.Ctx.WriteString(message)
	//	c.Ctx.Redirect(http.StatusFound, "/")
}
