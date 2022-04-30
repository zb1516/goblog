package until

import (
	"github.com/gorilla/sessions"
	"net/http"
)

//sessions 的存储库
var Store = sessions.NewCookieStore([]byte("session"))

// 当前会话
var Session *sessions.Session

// 用以获取会话
var Request *http.Request

// 用以写入会话
var Response http.ResponseWriter

// 初始化会话
func StartSession(w http.ResponseWriter, r *http.Request,name string) {
	session, _ := Store.Get(r, name)
	session.Save(r,w)
	Session = session
	Request = r
	Response = w
}

// Set 写入键值对应的会话数据
func Set(key string, value interface{}) {
	Session.Values[key] = value
	Save()
}

// Get 获取会话数据，获取数据时请做类型检测
func Get(key string) interface{} {
	return Session.Values[key]
}

// Destroy 删除某个会话项
func Destroy(key string) {
	delete(Session.Values, key)
	Save()
}

// Save 保持会话
func Save() {
	Session.Save(Request, Response)
}