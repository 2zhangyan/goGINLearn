package result

/**
封装返回消息体
*/
type Result struct {
	Code    int         `json:"code"`    //状态码
	Count   int         `json:"count"`   //总条数
	Message string      `json:"message"` // 消息
	Data    interface{} `json:"data"`    // 数据
}

/**
不带数据消息
*/
func ReturnNoData(code int, message string) *Result {
	return &Result{
		Code:    code,
		Message: message,
	}
}

/**
带数据的消息
*/
func ReturnData(code int, count int, message string, data interface{}) *Result {
	return &Result{
		Code:    code,
		Count:   count,
		Message: message,
		Data:    data,
	}
}
