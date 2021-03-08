package routers

import (
	"github.com/astaxie/beego"
	"zhiliao_blog/controllers/cms"
)

func init() {
    beego.Router("/cms",&cms.LoginController{})
    //beego.Router("/")
}
