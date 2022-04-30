package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

type handler func(w http.ResponseWriter,r *http.Request) (data interface{},err error)

func ResponseHandler(h handler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		//调用传入的方法
		data,err := h(w,r)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			log.Panic(err)
		}
		//处理响应结果，返回json格式
		errs := json.NewEncoder(w).Encode(&data)
		if errs != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			log.Panic(errs)
		}
	}
}