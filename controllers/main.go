package controllers

import (
	"beego_blog/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type  MainController  struct {
	beego.Controller
	Pager *models.Pager
}

//准备阶段，把接受当前页码。创建分页对象
func (c *MainController)Prepare(){
	var  page int
	var  err error
	//接受页码参数
	if   page  ,err = strconv.Atoi(c.Ctx.Input.Param(":page"));err !=nil{
		page =1
	}
	c.Pager = models.NewPager(page,2,0,"")
}


//显示首页
func  (c *MainController)  Index (){
	var list []*models.Post
	post :=models.Post{}
	//文章表的句柄
	//设置过滤条件,0正常，
	query := orm.NewOrm().QueryTable(&post).Filter("status", 0)
	//这里不查询所有了
	//_, err := query.All(&list)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(list)

	//查询符合条件的文章总数
	count, _ := query.Count()
	c.Pager.SetTotalnum(int(count))
	//设置urlpath的拼接规则
	c.Pager.SetUrlpath("/index%d.html")

	//严谨判断
	if  count>0 {
		//设置分页查询的偏移量
		//第1页   从索引0开始查询
		//第2页		从索引2开始
		//第3页    从索引4开始查询
		//偏移量		= (当前页码-1)* 每页大小
		offset :=(c.Pager.Page-1 ) * c.Pager.PageSize
		_, err := query.OrderBy("-istop","-views").Limit(c.Pager.PageSize, offset).All(&list)
		if err != nil{
			fmt.Println(err)
		}
	}
	c.Data["pagebar"] = c.Pager.PageBar()
	//list传到页面
	c.Data["list"]=list
	//Layout是指定不动的部分
	theme :="double"
	c.Layout = theme+"/layout.html"
	c.TplName = "double/index.html"
	//c.TplName = theme+"/index.html"

	//LayoutSections

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["head"]=theme+"/head.html"
	c.LayoutSections["banner"]=theme+"/banner.html"
	c.LayoutSections["middle"]=theme+"/middle.html"
	c.LayoutSections["right"]=theme+"/right.html"
	c.LayoutSections["foot"]=theme+"/foot.html"

}