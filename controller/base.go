package controller

import (
	"net/http"
	"text/template"
)

type ControllerInterface interface {
	Render()
}

type Controller struct {
	// 模版文件
	TplName        string
	ViewPath       string
	Layout         string
	LayoutSections map[string]interface{}
	TplPrefix      string
	TplExt         string
	EnableRender   bool
}

//渲染模版路径
func (c *Controller)Render(w http.ResponseWriter)  {
	//获取模版存放路径
	c.getTplPath()
	c.getTplExt()
	//要加载所有需要被嵌套的文件
	t, _ := template.ParseFiles(c.ViewPath+"/"+c.TplName+c.TplExt,c.ViewPath+"/layout/main"+c.TplExt,c.ViewPath+"/layout/header"+c.TplExt,c.ViewPath+"/layout/post_content"+c.TplExt,c.ViewPath+"/layout/content"+c.TplExt,c.ViewPath+"/layout/footer"+c.TplExt,c.ViewPath+"/layout/right"+c.TplExt)
	err:=t.ExecuteTemplate(w,c.TplName,c.LayoutSections)
	if err != nil {
		panic(err.Error())
	}
}

//模版存放路径
func (c *Controller)getTplPath() {
	if c.ViewPath == "" {
		c.ViewPath = "./view"
	}
}

//模版扩展名
func (c *Controller)getTplExt()  {
	if c.TplExt == "" {
		c.TplExt = ".html"
	}
}