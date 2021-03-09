package cms

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"zhiliao_blog/models"
	"zhiliao_blog/utils"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController)Get()  {
	l.TplName = "cms/login.html"
}

func (l *LoginController)Post()  {
	//TODO 特殊字符过滤 ？等字符 长度校验
	username :=l.GetString("username")
	password :=l.GetString("password")

	md5_pwd :=utils.GetMd5(password)

	o:=orm.NewOrm()
	exist := o.QueryTable(new(models.User)).Filter("user_name",username).Filter("password",md5_pwd).Exist()
	if exist{
		//跳转e
		l.SetSession("cms_user_name",username)
		l.Redirect(beego.URLFor("MainController.Get"),302)
		fmt.Println("登录成功")
	}else{
		l.Redirect(beego.URLFor("LoginController.Get"),302)
	}
}