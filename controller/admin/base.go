package admin

import (
	"goblog/until"
	"net/http"
	"reflect"
	"text/template"
)

type AdminInterface interface {
	Render()
	AuthLogin()
	Session()
}

type Admin struct {
	// 模版文件
	TplName        string
	ViewPath       string
	Layout         string
	LayoutContent  map[string]string
	LayoutSections map[string]interface{}
	TplPrefix      string
	TplExt         string
	EnableRender   bool
}

var UserId uint
var UserName string

//渲染模版路径
func (c *Admin)Render(w http.ResponseWriter)  {
	//获取模版存放路径
	c.getTplPath()
	c.getTplExt()
	//要加载所有需要被嵌套的文件
	t, _ := template.ParseFiles(c.ViewPath+"/"+c.Layout+"/"+c.TplName+c.TplExt,c.ViewPath+"/public/left"+c.TplExt,c.ViewPath+"/public/main"+c.TplExt,c.ViewPath+"/public/header"+c.TplExt)
	if len(c.LayoutContent) > 0 {
		err:=t.ExecuteTemplate(w,c.TplName,c.LayoutContent)
		if err != nil {
			panic(err.Error())
		}
	}else{
		err:=t.ExecuteTemplate(w,c.TplName,c.LayoutSections)
		if err != nil {
			panic(err.Error())
		}
	}
}

//模版存放路径
func (c *Admin)getTplPath() {
	if c.ViewPath == "" {
		c.ViewPath = "./view/admin"
	}
}

//模版扩展名
func (c *Admin)getTplExt()  {
	if c.TplExt == "" {
		c.TplExt = ".html"
	}
}

//验证用户是否登陆
func (c *Admin)AuthLogin(w http.ResponseWriter,r *http.Request) {
	//判断session是否存在，如果存在进行下一步，如果不存在，跳转到登陆页面
	until.StartSession(w,r,"user")	//开启session
	userValue:=until.Get("userData")
	if userValue == nil {
		http.Redirect(w,r,"/admin/login",http.StatusTemporaryRedirect)
		return
	}
	//通过反射获取结构体中的数据
	vf := reflect.ValueOf(userValue)
	userId := vf.FieldByName("ID")	//通过名字查询
	userName := vf.FieldByName("Username")	//通过名字查询
	UserId = uint(userId.Uint())	//把用户id转成uint类型
	UserName = userName.String()	//把用户名全局
}