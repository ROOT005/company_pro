package routers

import (
	"company_pro/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/submit", &controllers.SubmitController{})
	beego.Router("/haha", &controllers.LoginController{})
}
