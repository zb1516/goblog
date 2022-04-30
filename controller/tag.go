package controller

import (
	"github.com/julienschmidt/httprouter"
	"goblog/model"
	"goblog/service"
	"net/http"
	"strconv"
)

func TagList(w http.ResponseWriter,r *http.Request,_ httprouter.Params)  {
	var (
		pages string
		post model.Post
		tId int
		page int
		serv service.Service
		tagModel model.Tag
		tagPost model.TagPost
		controller Controller
		paginator service.Page
	)
	//接收分类参数
	tag:=r.URL.Query().Get("t")
	pages=r.URL.Query().Get("page")
	//判断参数是否存在
	if len(tag) > 0  {
		tId,_ = strconv.Atoi(tag)
	}else {
		tId = 0
	}
	if len(pages) >0 {
		page,_ = strconv.Atoi(pages)
	}else{
		page = 1
	}
	//查询出所有文章id
	tagPostIds,_:=tagPost.GetPostId(tId)
	//通过文章id查询出所有文章
	postRes ,total , postErr:=post.GetPostByTagPostId(page,8,tagPostIds)
	if postErr != nil {
		w.Write([]byte(postErr.Error()))
	}
	//从分类中取出分类名称
	tagName,_:=tagModel.GetTagName(tId)
	//右侧公共部分
	right:=serv.GetRight(w)
	//实例化分页
	paginate:=paginator.NewPage(page,8,total,"/tag?t="+tag)
	//把数据拼接到map中返回给前端页面
	data:=make(map[string]interface{})
	data["list"]=postRes
	data["newList"]=right["newList"]
	data["tag"]=right["tag"]
	data["category"]=right["category"]
	data["config"]=right["config"]
	data["tagName"]=tagName
	data["tId"] = tId
	data["catId"] = 0
	data["page"] = paginate.Show()			//渲染分页样式
	controller.TplName="tag"
	controller.LayoutSections=data
	controller.Render(w)
}