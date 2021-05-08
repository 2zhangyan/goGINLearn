package router

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
	"time"
)

//动态的路由
var (
	Router = gin.New()

	//限制GET请求次数
	GetLimiters = tollbooth.NewLimiter(100, &limiter.ExpirableOptions{
		ExpireJobInterval: time.Second,
	})

	//限制post请求次数
	PostLimiters = tollbooth.NewLimiter(100, &limiter.ExpirableOptions{
		ExpireJobInterval: time.Second,
	})

	v1 *gin.RouterGroup
)

func init() {
	v1 = Router.Group("v1")
	initRouter()
}

func initRouter() {
	initTest()
}
