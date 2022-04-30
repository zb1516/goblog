package controller

import (
	"github.com/julienschmidt/httprouter"
	"goblog/service"
	"net/http"
)

//关于我
func About(w http.ResponseWriter,r *http.Request , _ httprouter.Params) {
	var (
		controller Controller
		serv service.Service
		)
	//右侧公共部分
	right:=serv.GetRight(w)
	data := make(map[string]interface{})
	data["newList"]=right["newList"]
	data["tag"]=right["tag"]
	data["category"]=right["category"]
	data["config"]=right["config"]
	data["tId"] = 0
	data["catId"] = 0
	controller.TplName="about"
	controller.LayoutSections=data
	controller.Render(w)
}
