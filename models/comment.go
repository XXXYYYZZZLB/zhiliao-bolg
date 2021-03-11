package models

import "time"

type Comment struct {
	Id int `orm:"pl;auto"`
	Author *User `orm:"rel(fk);description(评论人)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime);description(创建时间)"`
	Content string `orm:"size(4000);description(评论内容)"`
	PId int `orm:"description(父级评论的id);default(0)"`
}

func(c *Comment) TableName()string{
	return "sys_post_comment"
}