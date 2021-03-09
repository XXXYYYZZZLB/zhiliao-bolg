package routers

import (
	"github.com/astaxie/beego"
	"zhiliao_blog/controllers/cms"
)

func init() {
    beego.Router("/cms",&cms.LoginController{})
	beego.Router("/cms/main/main",&cms.MainController{})
	beego.Router("/cms/main/welcome",&cms.MainController{},"get:Welcome")
    //beego.Router("/")
}
