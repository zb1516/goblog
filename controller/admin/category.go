package admin

import (
	"github.com/julienschmidt/httprouter"
	"goblog/model"
	"goblog/until"
	"net/http"
	"strconv"
)

//分类列表
func CategoryList(w http.ResponseWriter , r *http.Request , _ httprouter.Params)  {
	var (
		category model.Category
		admin Admin
	)
	admin.AuthLogin(w,r)	//验证用户是否登陆
	//查询标签列表
	categoryList , err:=category.GetCategoryList()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	data := make(map[string]interface{})
	data["list"] = categoryList
	admin.Layout = "category"
	admin.TplName="category"
	admin.LayoutSections=data
	admin.Render(w)
}

//添加分类
func CategoryAdd(w http.ResponseWriter , r *http.Request , _ httprouter.Params){
	var admin Admin
	admin.AuthLogin(w,r)	//验证用户是否登陆
	//如果是get请求加载页面
	admin.Layout = "category"
	admin.TplName="categoryAdd"
	admin.Render(w)
}

//保存分类
func CategorySave(w http.ResponseWriter,r *http.Request)(interface{},error)  {
	var (
		admin Admin
		category model.Category
		jsonResult until.Result
	)
	admin.AuthLogin(w,r)	//验证用户是否登陆
	r.ParseForm()
	category.Name =r.FormValue("name")
	//查询分类是否存在
	info , _ :=category.GetCategoryByName(category.Name)
	if len(info) > 0 {
		return jsonResult.Response(0,"分类已存在",nil) ,nil

	}
	//创建分类
	data,_:=category.CreateCategory(category)
	return jsonResult.Response(200,"添加成功",data) , nil
}


//修改分类
func CategoryEdit(w http.ResponseWriter , r *http.Request , _ httprouter.Params){
	var (
		admin Admin
		category model.Category
	)
	admin.AuthLogin(w,r)	//验证用户是否登陆
	//接受参数
	id , _ :=strconv.Atoi(r.URL.Query().Get("id"))
	//通过id查询分类是否存在
	cate , err:=category.GetCategoryById(id)
	if err != nil {
		panic(err)
	}
	data:=make(map[string]interface{})
	data["category"] = cate
	admin.LayoutSections=data
	admin.Layout = "category"
	admin.TplName="categoryEdit"
	admin.Render(w)
}

//保存编辑分类
func CategoryEditSave(w http.ResponseWriter,r *http.Request)(interface{},error)  {
	var (
		admin Admin
		category model.Category
		jsonResult until.Result
	)
	admin.AuthLogin(w,r)	//验证用户是否登陆
	r.ParseForm()
	category.Name =r.FormValue("name")
	cateId, _ := strconv.Atoi(r.FormValue("id"))
	category.ID = uint(cateId)
	//查询分类是否存在
	info , _ :=category.GetCategoryByName(category.Name)
	if len(info) > 0 {
		return jsonResult.Response(0,"分类已存在",nil) ,nil

	}
	//创建分类
	data,_:=category.UpdateCategory(category)
	return jsonResult.Response(200,"更新成功",data) , nil
}

//删除
func CategoryDel(w http.ResponseWriter,r *http.Request)(interface{},error)  {
	var (
		admin Admin
		category model.Category
		response until.Result
	)
	admin.AuthLogin(w,r)	//验证用户是否登陆
	r.ParseForm()
	cateId, _ := strconv.Atoi(r.FormValue("id"))
	category.ID = uint(cateId)
	//创建分类
	err:=category.DelCategory(category)
	if err != nil {
		return response.Response(200,"删除失败",nil) , nil
	}
	return response.Response(200,"删除成功",nil) , nil

}