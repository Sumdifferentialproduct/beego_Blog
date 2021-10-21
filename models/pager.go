package models

import (
	"bytes"
	"fmt"
)

//分页对象
type  Pager  struct {
	//当前页码
	Page		int
	//每页大小（每页多少文章）
	PageSize	int
	//文章总数量
	Totalnum	int
	//Url
	Urlpath  string
	//总页数,通过每页大小和文章总数计算
}

//分页的构造函数
func  NewPager(page  ,pageSize,totalnum int,urlpath string)*Pager{
	pager :=		new(Pager)
	pager.Page =page
	pager.PageSize = pageSize
	pager.Totalnum =	totalnum
	pager.Urlpath = urlpath
	return pager
}

//各种set方法
func(pager *Pager)SetPage(page int){
	pager.Page = page
}
func(pager *Pager)SetPageSize(pageSize int){
	pager.PageSize = pageSize
}
func(pager *Pager)SetTotalnum(totalnum int){
	pager.Totalnum = totalnum
}
func(pager *Pager)SetUrlpath(urlpath string){
	pager.Urlpath = urlpath
}

//拼接页面需要的pagebar
//需要显示10页的按钮，根据当前页码去计算
//拼接出的按钮，放到<a>标签就可以了
func (pager *Pager) PageBar()string{
	//查询总数量小于等于每页大小时，不需要分页导航
	if  pager.Totalnum<=pager.PageSize{
		return ""
	}
	//计算总页码
	var totalpage int
	//每页显示单数，除不尽，总数量/每页大小+1
	//每页显示双数。总数量/每页大小
	if pager.Totalnum%pager.PageSize!=0{
		totalpage = pager.Totalnum/pager.PageSize+1
	}else{
		totalpage = pager.Totalnum/pager.PageSize
	}
	//只展示10个分页导航按钮
	//申明需要显示的页码个数
	linknum := 10
	//从哪一页开始显示
	var  from int
	//显示到哪一页
	var to  int
	//偏移量
	offset:=5
	//总页码小于十的情况，从第一页显示到最后一页
	if   totalpage<linknum{
		from =10
		to =totalpage
	}else {
		//开始页码
		from = pager.Page - offset
		//最后一页页码
		to = from + linknum
		//判断特殊情况
		//判断起始页是负数的情况
		if from < 1 {
			from = 1
			to = from + linknum - 1
			//判断最后一页超过总页码的情况
		} else if to > totalpage {
			to = totalpage
			from = to - linknum + 1
		}
	}

		//用二进制类型拼字符串
		var  buf  bytes.Buffer
		buf.WriteString("<div class='page'>")
		//若当前页码是大于1时，显示下一页
		//Urlpath:="index%d.html"
		if   pager.Page>1{
			buf.WriteString(fmt.Sprintf("<a href='%s'>&raquo;</a>",
				fmt.Sprint(pager.Urlpath,pager.Page-1)))
		}
		//寻欢，根据from和to进行拼接
		for i :=from; i <=to;i++{
			//若是在当前页的
			if  i==pager.Page{
				buf.WriteString(fmt.Sprintf("<b>%d</b>",i))
			}else{
				//其他页是可点击的状态
				buf.WriteString(fmt.Sprintf("<a href='%s'>%d</a>",
					fmt.Sprintf(pager.Urlpath,i),i))
			}
		}

		//若当前页面小于总数时，显示下一页


	if  pager.Page<totalpage{
		buf.WriteString(fmt.Sprintf("<a href='%s'>&raquo;</a>",
			fmt.Sprintf(pager.Urlpath,pager.Page+1)))
	}
	//div结束标签
	buf.WriteString("</div>")
	//转string返回
	str := buf.String()
	return str
}
