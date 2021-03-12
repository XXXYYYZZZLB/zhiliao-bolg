package cms

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"zhiliao_blog/models"
)

type PostController struct {
	beego.Controller
}

func (p *PostController)Get(){
	o:=orm.NewOrm()
	posts := []models.Post{}
	o.QueryTable(new(models.Post)).RelatedSel().All(&posts)
	p.Data["posts"]=posts
	p.TplName="cms/post-list.html"
}
