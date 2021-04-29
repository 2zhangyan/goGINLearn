package router

import (
	"goGINLearn/handler"
)

func initTest() {
	test := v1.Group("test")

	test.GET("index", handler.LimitHandler(GetLimiters), handler.Index)

	//test.GET("/",func(context *gin.Context) {
	//	context.String(http.StatusOK, "hello gin")
	//})
}
