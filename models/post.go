package   models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"time"
)

//id
//userid
//author
//title
//color
//content
//tags
//views
//status
//posttime
//updated
//istop
//cover



type Post struct{
	Id   		int
	Userid		int
	Author		string		`orm:"size(15)"`
	Title		string		`orm:"size(100)"`
	Color		string		`orm:"size(7)"`
	Content		string		`orm:"type(text)"`
	Tags		string
	Views		int
	Status		int
	Posttime	time.Time	`orm:"type(datetime)"`
	Updated		time.Time	`orm:"type(datetime)"`
	Istop		int
	Cover		string
}

func   (post *Post)TableName()string{
	//从配置文件读取
	dbprefix := beego.AppConfig.String("dbprefix")
	return   dbprefix + "post"
}

//插入
func (post *Post) Insert()error{
	if _ ,err	:=orm.NewOrm().Insert(post);err !=nil{
		return err
	}
	return nil
}


//更新
func (post *Post) Update(fields ...string)error{
	if _ ,err	:=orm.NewOrm().Update(post,fields...);err !=nil{
		return err
	}
	return nil
}

//读取
func (post *Post) Read(fields ...string)error{
	if err	:=orm.NewOrm().Read(post,fields ...);err !=nil{
		return err
	}
	return nil
}


//Link显示文章链接
func (post *Post)	Link() string{
	//   /artice/id
	return	"/artice/"+strconv.Itoa(post.Id)

}

//ColorTitle  颜色标题
func (post *Post)ColorTitle() string{
	if post.Color !=""{
		return fmt.Sprintf("<span style='color:%s'>%s</span>",post.Color,post.Title)
	}
	return  post.Title

}

//Excerpt  文章内容  表里有这个字段  以字符串返回
func (post *Post)Excerpt() string{
	return post.Content

}

//TagLink  所属分类  文章对应的标签
func (post *Post) TagsLink() string {
	if post.Tags == "" {
		return ""
	}
	// 去掉字符串的首尾逗号
	tagslink := strings.Trim(post.Tags, ",")
	return tagslink
}

//删除文档的方法和前两个不太一样.后面重写，需要多表操作