package admin

import (
	"github.com/julienschmidt/httprouter"
	"goblog/model"
	"goblog/until"
	"net/http"
	"strconv"
)
//文章列表
func PostList(w http.ResponseWriter , r *http.Request , _ httprouter.Params)  {
	var (
		post model.Post
		admin Admin
	)
	admin.AuthLogin(w,r)	//验证用户是否登陆
	//查询文章列表
	postList,total , err:=post.GetPostList(0,0,0)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	data := make(map[string]interface{})
	data["list"] = postList
	data["total"] = total
	admin.Layout = "post"
	admin.TplName="post"
	admin.LayoutSections=data
	admin.Render(w)
}

//发布文章
func PostAdd(w http.ResponseWriter , r *http.Request , _ httprouter.Params){
	var (
		admin Admin
		tag model.Tag
		category model.Category
	)
	admin.AuthLogin(w,r)	//验证用户是否登陆
	//获取所有标签
	tagList , _:=tag.GetTagList()
	//获取所有分类
	categoryList,_:=category.GetCategoryList()
	data:=make(map[string]interface{})
	data["taglist"] = tagList
	data["catelist"] = categoryList
	//如果是get请求加载页面
	admin.Layout = "post"
	admin.TplName="postAdd"
	admin.LayoutSections=data
	admin.Render(w)
}

//保存文章
func PostSave(w http.ResponseWriter,r *http.Request) (interface{},error) {
	var (
		response until.Result
		post model.Post
	)
	//判断请求方式
	if r.Method == "GET" {
		return response.Response(0,"错误的请求方式",nil),nil
	}
	r.ParseForm()
	post.Title=r.FormValue("title")
	post.Url=r.FormValue("url")
	post.Content=r.FormValue("content")
	post.Tags=r.FormValue("tags")
	post.IsTop,_ = strconv.Atoi(r.FormValue("is_top"))
	post.CategoryId , _= strconv.Atoi(r.FormValue("category_id"))
	post.Image=r.FormValue("image")
	post.Info=r.FormValue("info")
	post.UserId=int(UserId)		//获取用户id
	post.Username=UserName		//获取用户名
	//创建文章
	data,err:=post.CreatePost(post)
	if err != nil {
		return response.Response(0,"保存失败",nil),err
	}
	//接受参数
	return response.Response(200,"保存成功",data),err
}


//修改文章
func PostEdit(w http.ResponseWriter , r *http.Request , _ httprouter.Params){
	var (
		admin Admin
		post model.Post
		tag model.Tag
		category model.Category
	)
	admin.AuthLogin(w,r)	//验证用户是否登陆
	//接受参数
	id , _ :=strconv.Atoi(r.URL.Query().Get("id"))
	//通过id查询是否存在
	post , err:=post.GetPostById(id)
	if err != nil {
		panic(err)
	}
	//获取所有标签
	tagList , _:=tag.GetTagList()
	//获取所有分类
	categoryList,_:=category.GetCategoryList()
	data:=make(map[string]interface{})
	data["post"] = post
	data["taglist"] = tagList
	data["catelist"] = categoryList
	admin.LayoutSections=data
	admin.Layout = "post"
	admin.TplName="postEdit"
	admin.Render(w)
}

//保存编辑文章
func PostEditSave(w http.ResponseWriter,r *http.Request)(interface{},error)  {
	var (
		admin Admin
		post model.Post
		jsonResult until.Result
	)
	admin.AuthLogin(w,r)	//验证用户是否登陆
	r.ParseForm()
	post.Title=r.FormValue("title")
	post.Url=r.FormValue("url")
	post.Content=r.FormValue("content")
	post.Tags=r.FormValue("tags")
	post.IsTop,_ = strconv.Atoi(r.FormValue("is_top"))
	post.CategoryId , _= strconv.Atoi(r.FormValue("category_id"))
	post.Image=r.FormValue("image")
	post.Info=r.FormValue("info")
	post.UserId=int(UserId)		//获取用户id
	post.Username=UserName		//获取用户名
	postId, _ := strconv.Atoi(r.FormValue("id"))
	post.ID = uint(postId)
	//更新文章
	data,_:=post.UpdatePost(post)
	return jsonResult.Response(200,"更新成功",data) , nil
}

//删除
func PostDel(w http.ResponseWriter,r *http.Request)(interface{},error)  {
	var (
		admin Admin
		post model.Post
		response until.Result
	)
	admin.AuthLogin(w,r)	//验证用户是否登陆
	r.ParseForm()
	postId, _ := strconv.Atoi(r.FormValue("id"))
	post.ID = uint(postId)
	//创建分类
	err:=post.DelPost(post)
	if err != nil {
		return response.Response(200,"删除失败",nil) , nil
	}
	return response.Response(200,"删除成功",nil) , nil

}