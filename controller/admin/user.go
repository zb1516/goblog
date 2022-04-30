package admin

import (
	"encoding/gob"
	"github.com/julienschmidt/httprouter"
	"goblog/model"
	"goblog/until"
	"log"
	"net/http"
	"strconv"
	"time"
)

//获取用户列表
func UserList (w http.ResponseWriter,r *http.Request,_ httprouter.Params) {
	if r.Method == "post" {
		w.Write([]byte("非法请求！"))
	}
	var (
		user model.User
		admin Admin
	)
	admin.AuthLogin(w,r)	//验证用户是否登陆
	//如果是get方式请求
	res , err := user.GetUserlist()
	if err != nil {
		log.Println(err.Error())
	}
	data := make(map[string]interface{})
	data["list"] = res
	admin.Layout="user"
	admin.TplName="user"
	admin.LayoutSections=data
	admin.Render(w)
}

//渲染添加页面
func UserAdd(w http.ResponseWriter,r *http.Request,_ httprouter.Params) {
	var (
		admin Admin
	)
	admin.AuthLogin(w,r)	//验证用户是否登陆
	//添加用户
	data := make(map[string]interface{})
	admin.Layout="user"
	admin.TplName="userAdd"
	admin.LayoutSections=data
	admin.Render(w)
}

//保存用户信息
func UserSave(w http.ResponseWriter,r *http.Request)(interface{},error)  {
	var (
		admin Admin
		user model.User
		jsonResult until.Result
	)
	admin.AuthLogin(w,r)	//验证用户是否登陆
	r.ParseForm()
	user.Username =r.FormValue("username")
	user.Password,_ = until.Md5(r.FormValue("password"))		//md5加密
	user.Email =r.FormValue("email")
	//查询用户是否存在
	info , err :=user.GetUserByName(user.Username)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	if len(info) > 0 {
		return jsonResult.Response(0,"用户已存在",nil) ,nil

	}
	//如果用户不存在的时候创建用户
	u,_:=user.CreateUser(user)
	return jsonResult.Response(200,"添加成功",u) , nil
}

//用户登陆
func Login(w http.ResponseWriter,r *http.Request , _ httprouter.Params) {
	var (
		admin Admin
		)
	data := make(map[string]interface{})
	admin.Layout="user"
	admin.TplName="login"
	admin.LayoutSections=data
	admin.Render(w)
}

//登陆执行方法
func Logins(w http.ResponseWriter,r *http.Request) (interface{},error){
	var (
		user model.User
		response until.Result
	)
	//判断请求方式
	if r.Method == "GET" {
		w.Write([]byte("错误的请求方式"))
	}
	//接受参数
	r.ParseForm()
	username := r.FormValue("username")
	password,_ := until.Md5(r.FormValue("password"))
	//查询用户信息
	userData , err:=user.GetUser(username,password)
	if err != nil {
		return response.Response(0,"用户不存在",userData),err
	}
	//更新数据结构
	update:=make(map[string]interface{})
	update["last_ip"] , _ = until.GetLocalIP()
	update["last_time"] = time.Now()
	userRes,err:=user.UpdateUser(username,update)
	gob.Register(userRes)
	until.StartSession(w,r,"user")
	until.Set("userData",userRes)
	return response.Response(200,"登陆成功",userRes),err
}

//退出登陆
func Logout(w http.ResponseWriter,r *http.Request,_ httprouter.Params)  {
	until.StartSession(w,r,"user")
	until.Destroy("user")
	http.Redirect(w,r,"/admin/login",http.StatusTemporaryRedirect)
}


//修改用户
func UserEdit(w http.ResponseWriter , r *http.Request , _ httprouter.Params){
	var (
		admin Admin
		user model.User
	)
	admin.AuthLogin(w,r)	//验证用户是否登陆
	//接受参数
	id , _ :=strconv.Atoi(r.URL.Query().Get("id"))
	//通过id查询是否存在
	user , err:=user.GetUserById(id)
	if err != nil {
		panic(err)
	}
	data:=make(map[string]interface{})
	data["user"] = user
	admin.LayoutSections=data
	admin.Layout="user"
	admin.TplName="userEdit"
	admin.Render(w)
}

//保存编辑用户
func UserEditSave(w http.ResponseWriter,r *http.Request)(interface{},error)  {
	var (
		admin Admin
		user model.User
		jsonResult until.Result
	)
	admin.AuthLogin(w,r)	//验证用户是否登陆
	r.ParseForm()
	user.Username =r.FormValue("username")
	user.Password,_= until.Md5(r.FormValue("password"))
	user.Email =r.FormValue("email")
	userId, _ := strconv.Atoi(r.FormValue("id"))
	user.ID = uint(userId)
	//查询用户是否存在
	info , _ :=user.GetUserByName(user.Username)
	if len(info) > 0 {
		return jsonResult.Response(0,"用户名已存在",nil) ,nil
	}
	//创建分类
	err :=user.UpdateUserById(user)
	if err != nil {
		return jsonResult.Response(0,"更新失败",nil) , err
	}
	return jsonResult.Response(200,"更新成功",nil) , nil
}

//删除
func UserDel(w http.ResponseWriter,r *http.Request)(interface{},error)  {
	var (
		admin Admin
		user model.User
		response until.Result
	)
	admin.AuthLogin(w,r)	//验证用户是否登陆
	r.ParseForm()
	cateId, _ := strconv.Atoi(r.FormValue("id"))
	user.ID = uint(cateId)
	//创建分类
	err:=user.DelUser(user)
	if err != nil {
		return response.Response(200,"删除失败",nil) , nil
	}
	return response.Response(200,"删除成功",nil) , nil
}