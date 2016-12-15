package controllers

import (
	"company_pro/models"
	"github.com/astaxie/beego"
)

type SubmitController struct {
	beego.Controller
}

func (c *SubmitController) Get() {
	c.Redirect("/", 301)
	return
}
func (c *SubmitController) Post() {
	name := c.Input().Get("name")
	account := c.Input().Get("account")
	phone := c.Input().Get("phonenum")
	//varify := c.Input().Get("verify")
	err := models.AddUser(name, account, phone)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/", 302)
	return
}
