package admin

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"goblog/model"
	"goblog/until"
	"net/http"
)

type AdminConfig struct {
	Admin
}

func Config(w http.ResponseWriter,r *http.Request,_ httprouter.Params)  {
	controller:=&AdminConfig{}
	var (
		config model.Config
	)
	//查询配置信息
	configRes , err :=config.GetConfig()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	data:=make(map[string]string)
	if len(configRes) > 0 {
		for _,val := range configRes {
			data[val.Name]=val.Value
		}
	}
	controller.LayoutContent=data
	controller.Layout="config"
	controller.TplName="config"
	controller.Render(w)
}

//保存配置信息
func ConfigSave(w http.ResponseWriter,r *http.Request) (interface{},error) {
	configParam:=&model.CofnigParam{}
	var (
		config model.Config
		response until.Result
	)
	r.ParseForm()
	configParam.Title = r.FormValue("title")
	configParam.Url = r.FormValue("url")
	configParam.Keywords = r.FormValue("keywords")
	configParam.Description = r.FormValue("description")
	configParam.Email = r.FormValue("email")
	configParam.Timezone = r.FormValue("timezone")
	configParam.Start = r.FormValue("start")
	configParam.QQ = r.FormValue("qq")
	configParam.Github = r.FormValue("github")
	configParam.Csdn = r.FormValue("csdn")
	b, _ := json.Marshal(&configParam)
	data:=make(map[string]string)
	_ = json.Unmarshal(b, &data)
	res:=config.UpdateConfig(data)
	if !res {
		return response.Response(0,"保存失败",nil),nil
	}
	return response.Response(200,"保存成功",configParam),nil
}