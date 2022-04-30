package controller

import (
	"github.com/julienschmidt/httprouter"
	"goblog/model"
	"goblog/service"
	"net/http"
	"strconv"
)

func CategoryList(w http.ResponseWriter,r *http.Request , _ httprouter.Params)  {
	var (
		pages string
		post model.Post
		catId int
		page int
		serv service.Service
		category model.Category
		controller Controller
		paginator service.Page
	)
	//接收分类参数
	c:=r.URL.Query().Get("c")
	pages=r.URL.Query().Get("page")
	//判断参数是否存在
	if len(c) > 0  {
		catId,_ = strconv.Atoi(c)
	}else{
		catId=0
	}
	if len(pages) >0 {
		page,_ = strconv.Atoi(pages)
	}else{
		page = 1
	}
	//查询出分类下文章
	postRes ,total, postErr:=post.GetPostByCategory(page,8,catId)
	if postErr != nil {
		w.Write([]byte(postErr.Error()))
	}
	//从分类中取出分类名称
	categoryName,_:=category.GetCategoryName(catId)
	//右侧公共部分
	right:=serv.GetRight(w)
	//实例化分页
	paginate:=paginator.NewPage(page,8,total,"/category?c="+c)
	//把数据拼接到map中返回给前端页面
	data:=make(map[string]interface{})
	data["list"]=postRes
	data["newList"]=right["newList"]
	data["category"]=right["category"]
	data["tag"]=right["tag"]
	data["config"]=right["config"]
	data["catId"] = catId
	data["tId"] = 0
	data["tagName"]=categoryName
	data["page"] = paginate.Show()			//渲染分页样式
	controller.TplName="category"
	controller.LayoutSections=data
	controller.Render(w)
}