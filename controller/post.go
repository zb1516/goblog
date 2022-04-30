package controller

import (
	"github.com/julienschmidt/httprouter"
	"goblog/model"
	"goblog/until"
	"net/http"
	"strconv"
)

//文章详情
func Post(w http.ResponseWriter,r *http.Request , _ httprouter.Params) {
	//判断请求方式
	if r.Method == "post" {
		w.Write([]byte("非法请求！"))
	}
	var (
		id int
		post model.Post
		controller Controller
	)
	//如果是get请求
	postId := r.URL.Query().Get("id")
	id,_ = strconv.Atoi(postId)
	res,err:=post.GetPostInfo(id)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	data:=make(map[string]interface{})
	data["postInfo"] = res
	data["is_post"] = true	//文章内容标示
	controller.TplName="post"
	controller.LayoutSections=data
	controller.Render(w)
}

//喜欢接口
func Love(w http.ResponseWriter,r *http.Request) (interface{},error) {
	//判断是否是post提交
	if r.Method == "get" {
		w.Write([]byte("非法请求!"))
	}
	var (
		post model.Post
		love int
		response until.Result
	)
	r.ParseForm()
	postId,_ := strconv.Atoi(r.FormValue("id"))
	//对文章喜欢数+1
	postRes,err:=post.PostLoveUpate(postId)
	if err != nil {
		return response.Response(0,"",nil),err
	}
	for _,v := range postRes {
		love = v.Love
	}
	return response.Response(200,"",love),nil
}