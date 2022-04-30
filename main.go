package main

import (
	"goblog/config"
	"goblog/router"
	"net/http"
	_ "net/http/pprof"
)

func main()  {
	//加载配置
	config.LoadConfig()
	//注册路由
	router:=router.RegisterRouter()
	http.ListenAndServe("127.0.0.1:8080",router)
}