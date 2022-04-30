package until

type Result struct {
	Code int         `json:"code"` // 错误码
	Message  string  `json:"message"`  // 错误描述
	Data interface{} `json:"data"` // 返回数据
}

//直接返回响应数据
func (d *Result)Response(code int , message string , data interface{}) *Result {
	return &Result{
		Code: code,
		Message: message,
		Data: data,
	}
}