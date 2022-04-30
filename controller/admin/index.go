package admin

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Index(w http.ResponseWriter , r *http.Request , _ httprouter.Params)  {
	var (
		admin Admin
	)
	admin.AuthLogin(w,r)	//验证用户是否登陆
	admin.Layout="index"
	admin.TplName="index"
	admin.Render(w)
}
