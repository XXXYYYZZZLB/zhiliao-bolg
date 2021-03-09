package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "zhiliao_blog/routers"
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	_"zhiliao_blog/models"
	"zhiliao_blog/utils"
)

func init()  {
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	database := beego.AppConfig.String("database")

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local",username,password,host,port,database)
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql",datasource)

	//迁移表
	fmt.Println("连接数据库")
	name := "default"
	force := false
	verbose := true
	err := orm.RunSyncdb(name,force,verbose)
	if err != nil{
		panic(err)
	}
}

func main() {

	beego.InsertFilter("/cms/main/*",beego.BeforeRouter,utils.CmsLoginFilter)
	orm.RunCommand()
	beego.Run()
}

