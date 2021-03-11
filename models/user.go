package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id int `orm:"pk;auto"`
	UserName string `orm:"description(用户名);index;unique`
	Password string `orm:"description(密码)"`
	IsAdmin int `orm:"description(1是管理员，2是普通用户);default(2)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime);description(创建时间)"`
	Cover string `orm:"description(用户图像);default(static/upload/bq3.png)"`

	Post []*Post `orm:"reverse(many)"`
	Comment []*Comment `orm:"reverse(many)"`
}

func (u *User)TableName() string {//表名
	return "sys_user"
}

func init()  {//注册
	orm.RegisterModel(new(User),new(Post),new(Comment))
}
