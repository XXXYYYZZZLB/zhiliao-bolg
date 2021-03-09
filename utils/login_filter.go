package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

//登录限制
func CmsLoginFilter(ctx *context.Context)  {
	cms_user_name := ctx.Input.Session("cms_user_name")

	if cms_user_name == nil {
		ctx.Redirect(302,beego.URLFor("LoginController.Get"))
	}
}
