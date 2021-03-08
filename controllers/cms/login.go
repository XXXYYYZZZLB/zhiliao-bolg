package cms

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (l *LoginController)Get()  {
	l.TplName = "cms/login.html"
}

func (l *LoginController)Post()  {
	
}