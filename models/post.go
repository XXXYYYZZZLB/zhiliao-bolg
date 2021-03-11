package models

import (
	"time"
)

//type Tag{
//	Id int
//	Name string
//}

//帖子模型
type Post struct {
	Id int `orm:"pl;auto"`
	Title string `orm:"description(标题)"`
	Desc string `orm:"description(描述)"`
	Cover string  `orm:"description(封面图);default(static/upload/bq3.png)"`
	ReadNum int `orm:"description(阅读数);default(0)"`
	StarNum int `orm:"description(点赞数);default(0)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime);description(创建时间)"`

	Author *User `orm:"description(作者);rel(fk)"`

}

func (p *Post)TableName() string  {
	return "sys_post"
}
