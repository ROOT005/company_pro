package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}
func (c *LoginController) Post() {
	uname := stringTomd5(c.Input().Get("uname"), 32)
	pwd := stringTomd5(c.Input().Get("pwd"), 32)
	if beego.AppConfig.String("uname") == stringTomd5(uname, 32) && beego.AppConfig.String("pwd") == stringTomd5(pwd, 32) {
		maxAge := 1<<31 - 1
		c.Ctx.SetCookie("uname", uname, maxAge, "/")
		c.Ctx.SetCookie("pwd", pwd, maxAge, "/")
		c.Redirect("/", 303)
	}
	c.Redirect("/", 303)
}

//md5加密封装
func stringTomd5(str string, size int8) string {
	sourcedata := []byte(str)
	hash := md5.New()
	hash.Write(sourcedata)
	cipherText := hash.Sum(nil)
	hexText := make([]byte, size)
	hex.Encode(hexText, cipherText)
	return string(hexText)
}

//检查是否登录封装
func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value
	return beego.AppConfig.String("uname") == stringTomd5(uname, 32) && beego.AppConfig.String("pwd") == stringTomd5(pwd, 32)
}
