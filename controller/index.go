package controller

import (
	"github.com/julienschmidt/httprouter"
	"goblog/model"
	"goblog/service"
	"net/http"
	"strconv"
)

//获取首页推荐列表
func Index (w http.ResponseWriter,r *http.Request, _ httprouter.Params) {
	if r.Method == "post" {
		w.Write([]byte("非法请求！"))
	}
	var (
		page int
		post model.Post
		serv service.Service
		controller Controller
		paginator service.Page
	)
	pagestr := r.URL.Query().Get("page") //当前页
	if len(pagestr) > 0 {
		page,_ = strconv.Atoi(pagestr)
	}else{
		page = 1
	}
	//如果是get方式请求
	rows ,total , err := post.GetPostList(page,10,1)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	//实例化分页
	paginate:=paginator.NewPage(page,10,total,"/")
	//创建分页
	//右侧公共模块
	right:=serv.GetRight(w)
	data := make(map[string]interface{})
	data["list"]=rows
	data["newList"]=right["newList"]
	data["tag"]=right["tag"]
	data["category"]=right["category"]
	data["tId"] = 0
	data["catId"] = 0
	data["config"]=right["config"]
	data["page"] = paginate.Show()			//渲染分页样式
	controller.TplName="index"
	controller.LayoutSections=data
	controller.Render(w)
}
