package router

import (
	"goGINLearn/handler"
)

func initTest() {
	test := v1.Group("test")

	test.GET("index", handler.LimitHandler(GetLimiters), handler.Index)
	test.POST("post", handler.LimitHandler(PostLimiters), handler.Post)
	test.GET("get", handler.LimitHandler(PostLimiters), handler.Get)
	test.POST("query-post", handler.LimitHandler(PostLimiters), handler.QueryPost)
	test.POST("map", handler.LimitHandler(PostLimiters), handler.Map)

	//test.GET("/",func(context *gin.Context) {
	//	context.String(http.StatusOK, "hello gin")
	//})
}
