//package models
//
//import (
//	_ "Blog/routers"
//	"fmt"
//	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/orm"
//	_ "github.com/go-sql-driver/mysql"
//)
//
//func  init (){
//
//	//dbhost = 127.0.0.1
//	//dbport = 3306
//	//dbuser = root
//	//dbpassword = 123
//	//dbname = beego_blog
//
//
//	//从配置文件获取配置信息
//	dbhost := beego.AppConfig.String("dbhost")
//	dbport := beego.AppConfig.String("dbport")
//	dbuser := beego.AppConfig.String("dbuser")
//	dbpassword := beego.AppConfig.String("dbpassword")
//	dbname := beego.AppConfig.String("dbname")
////拼接数据库连接
//	dburl :=dbuser+":"+dbpassword+"@tcp("+dbhost+":"+dbport+")/"+dbname+"?charset=utf8"
//	fmt.Println("--------------",dburl,"-------------")
//
//	//1.注册库
//	err := orm.RegisterDataBase("default", "mysql",
//		dburl,50,20)
//	if err != nil{
//		fmt.Println("mysql连接失败",err)
//		return
//	}
//
//	//注册model表
//	orm.RegisterModel(new(Link),new(Post),new(Mood),new(Tag),new(TagPost),new(User))
//	//设置打印sql
//	orm.Debug =  true
//}


package models

import (
	"fmt"
	// 加载mysql驱动
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//dbhost = 127.0.0.1
//dbport = 3306
//dbuser = root
//dbpassword = admin
//dbname = beego_blog
func init() {
	// 从配置文件获取配置信息
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	// 拼接数据库连接
	// "root:admin@tcp(localhost:3306)/beego_blog?charset=utf8"
	// beego orm不需要 parseTime=true
	dburl := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	fmt.Println("dburl:---" + dburl)
	// 设置default数据库
	err := orm.RegisterDataBase("default", "mysql", dburl, 50, 20)
	if err != nil {
		fmt.Println("mysql连接失败：", err)
	}
	// 注册model
	orm.RegisterModel(new(Link), new(Mood), new(Post), new(Tag), new(TagPost), new(User))
	// 设置打印SQL
	orm.Debug = true
}
