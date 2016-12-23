package controllers

import (
	"company_pro/models"
	//"fmt"
	"github.com/astaxie/beego"
	//"strconv"
)

type SubmitController struct {
	beego.Controller
}

func (c *SubmitController) Get() {
	c.Redirect("/", 301)
	return
}
func (c *SubmitController) Post() {
	//id, value := c.GetString("captcha_id"), c.GetString("captcha")
	//b := captcha.Verify(id, value) //验证码校验
	name := c.Input().Get("name")
	account := c.Input().Get("account")
	phone := c.Input().Get("phonenum")
	err := models.AddUser(name, account, phone)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/", 302)
	return
}
