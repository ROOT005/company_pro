package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"strconv"
)

type VerifyController struct {
	beego.Controller
}

var cpt *captcha.Captcha

func init() {
	//验证码
	var store = cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 4
	cpt.StdWidth = 100
	cpt.StdHeight = 34
}
func (c *VerifyController) Get() {
	//验证码
	result := cpt.Verify(c.Input().Get("captcha_id"), c.Input().Get("id"))
	c.Ctx.WriteString(strconv.FormatBool(result))
}
