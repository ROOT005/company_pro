package routers

import (
	"company_pro/controllers"
	"company_pro/models"
	"github.com/astaxie/beego"
)

func init() {
	//注册路由
	models.RegisterDB()
	//设置路由
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/submit", &controllers.SubmitController{})
}
