package admin

import (
	"github.com/julienschmidt/httprouter"
	"goblog/model"
	"goblog/until"
	"net/http"
	"strconv"
)

//标签列表
func TagList(w http.ResponseWriter , r *http.Request , _ httprouter.Params)  {
	var (
		tag model.Tag
		admin Admin
	)
	admin.AuthLogin(w,r)	//验证用户是否登陆
	//查询标签列表
	tagList , err:=tag.GetTagList()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	data := make(map[string]interface{})
	data["list"] = tagList
	admin.Layout="tag"
	admin.TplName="tag"
	admin.LayoutSections=data
	admin.Render(w)
}

//添加标签
func TagAdd(w http.ResponseWriter , r *http.Request , _ httprouter.Params){
	var admin Admin
	admin.AuthLogin(w,r)	//验证用户是否登陆
	//如果是get请求加载页面
	admin.Layout="tag"
	admin.TplName="tagAdd"
	admin.Render(w)
}

//保存标签
func TagSave(w http.ResponseWriter,r *http.Request)(interface{},error)  {
	var (
		admin Admin
		tag model.Tag
		jsonResult until.Result
	)
	admin.AuthLogin(w,r)	//验证用户是否登陆
	r.ParseForm()
	tag.Name =r.FormValue("name")
	//查询标签是否存在
	info , err :=tag.GetTagByName(tag.Name)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	if len(info) > 0 {
		return jsonResult.Response(0,"标签已存在",nil) ,nil

	}
	//创建标签
	data,_:=tag.CreateTag(tag)
	return jsonResult.Response(200,"添加成功",data) , nil
}

//修改标签
func TagEdit(w http.ResponseWriter , r *http.Request , _ httprouter.Params){
	var (
		admin Admin
		tag model.Tag
	)
	admin.AuthLogin(w,r)	//验证用户是否登陆
	//接受参数
	id , _ :=strconv.Atoi(r.URL.Query().Get("id"))
	//通过id查询分类是否存在
	cate , err:=tag.GetTagById(id)
	if err != nil {
		panic(err)
	}
	data:=make(map[string]interface{})
	data["tag"] = cate
	admin.LayoutSections=data
	admin.Layout="tag"
	admin.TplName="tagEdit"
	admin.Render(w)
}

//保存编辑标签
func TagEditSave(w http.ResponseWriter,r *http.Request)(interface{},error)  {
	var (
		admin Admin
		tag model.Tag
		jsonResult until.Result
	)
	admin.AuthLogin(w,r)	//验证用户是否登陆
	r.ParseForm()
	tag.Name =r.FormValue("name")
	tagId, _ := strconv.Atoi(r.FormValue("id"))
	tag.ID = uint(tagId)
	//查询分类是否存在
	info , _ :=tag.GetTagByName(tag.Name)
	if len(info) > 0 {
		return jsonResult.Response(0,"标签已存在",nil) ,nil

	}
	//创建分类
	data,_:=tag.UpdateTag(tag)
	return jsonResult.Response(200,"更新成功",data) , nil
}

//删除
func TagDel(w http.ResponseWriter,r *http.Request)(interface{},error)  {
	var (
		admin Admin
		tag model.Tag
		response until.Result
	)
	admin.AuthLogin(w,r)	//验证用户是否登陆
	r.ParseForm()
	tagId, _ := strconv.Atoi(r.FormValue("id"))
	tag.ID = uint(tagId)
	//创建分类
	err:=tag.DelTag(tag)
	if err != nil {
		return response.Response(200,"删除失败",nil) , nil
	}
	return response.Response(200,"删除成功",nil) , nil

}