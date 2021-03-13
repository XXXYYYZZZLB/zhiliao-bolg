package cms

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"zhiliao_blog/models"
	"zhiliao_blog/utils"
)

type PostController struct {
	beego.Controller
}

func (p *PostController) Get()  {

	o := orm.NewOrm()
	posts := []models.Post{}

	qs := o.QueryTable(new(models.Post))


	qs.RelatedSel().All(&posts)


	count,_ := qs.Count()

	current_page,err := p.GetInt("p")
	if err != nil {
		current_page = 1
	}

	page_size := 10

	total_pages := utils.GetPageNum(count,page_size)

	// 前后页码
	arround_count := 4
	left_pages, right_pages, left_has_more, right_has_more := utils.Get_pagination_data(total_pages,current_page,arround_count)


	has_pre_page, has_next_page, pre_page, next_page := utils.HasNext(current_page, total_pages)

	// 100 , 100
	fmt.Println(left_pages)
	fmt.Println(right_pages)
	fmt.Println(current_page)
	fmt.Println(left_has_more)
	fmt.Println(right_has_more)
	fmt.Println(left_pages)
	p.Data["left_pages"] = left_pages
	p.Data["left_has_more"] = left_has_more
	p.Data["page"] = current_page

	p.Data["has_pre_page"] = has_pre_page
	p.Data["pre_page"] = pre_page
	p.Data["has_next_page"] = has_next_page
	p.Data["next_page"] = next_page

	p.Data["right_pages"] = right_pages
	p.Data["right_has_more"] = right_has_more


	p.Data["num_pages"] = total_pages //总页数
	p.Data["count"] = count //总数量
	p.Data["posts"] = posts
	p.TplName = "cms/post-list.html"

}
