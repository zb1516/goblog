package controller

import (
	"github.com/julienschmidt/httprouter"
	"goblog/until"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

//上传文件
func Upload(rw http.ResponseWriter, r *http.Request) (interface{},error)  {
	var (
		response until.Result
	)
	//接收上传文件
	file, handler, err := r.FormFile("file")
	if err != nil {
		return response.Response(0,"上传失败",nil),err
	}
	defer file.Close()
	filePath:="./upload/"+handler.Filename
	//创建上传路径
	os.Mkdir("./upload",os.ModePerm)
	f,err :=os.Create(filePath)
	if err != nil {
		return response.Response(0,"上传失败",nil),err
	}
	defer f.Close()
	io.Copy(f, file)
	//创建map
	data:=make(map[string]interface{})
	data["file"]=map[string]interface{}{
		"path":"/upload/"+handler.Filename,
	}
	//返回给前端
	return response.Response(200,"上传成功",data),err
}

// 显示图片接口
func ShowPicHandle( w http.ResponseWriter, r *http.Request , _ httprouter.Params) {
	path := r.URL.Query().Get("path")
	file, err := os.Open("." + path)
	errorHandle(err, w)

	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	errorHandle(err, w)
	w.Write(buff)
}

// 统一错误输出接口
func errorHandle(err error, w http.ResponseWriter) {
	if  err != nil {
		w.Write([]byte(err.Error()))
	}
}
