package controllers

import (
	"company_pro/models"
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
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
func (c *SubmitController) PhoneNum() {
	phonenum := c.Input().Get("phonenum")
	valid := validation.Validation{}
	if valid.Phone(phonenum, "Mobile").Ok && valid.Required(phonenum, "Mobile").Ok {
		c.Ctx.WriteString("true")
	} else {
		c.Ctx.WriteString("false")
	}
}
