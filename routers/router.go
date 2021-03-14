package routers

import (
	"github.com/astaxie/beego"
	"zhiliao_blog/controllers/cms"
)

func init() {
    beego.Router("/cms",&cms.LoginController{})
	beego.Router("/cms/main/main",&cms.MainController{})
	beego.Router("/cms/main/welcome",&cms.MainController{},"get:Welcome")
	beego.Router("/cms/main/post_list",&cms.PostController{})
	beego.Router("/cms/main/post_toadd",&cms.PostController{},"get:ToAdd")
	beego.Router("/cms/main/post_doadd",&cms.PostController{},"post:DoAdd")

}
