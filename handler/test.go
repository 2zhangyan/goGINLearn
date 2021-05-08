package handler

import (
	"github.com/gin-gonic/gin"
	"goGINLearn/model/test"
	"net/http"
)

//get
func Index(g *gin.Context) {
	g.JSON(http.StatusOK, test.Index())
}

//post
func Post(g *gin.Context) {
	params := make(map[string]string)
	params["name"] = g.PostForm("name")
	params["age"] = g.PostForm("age")
	g.JSON(http.StatusOK, test.Post(params))
}

//获取Query参数
func Get(g *gin.Context) {
	params := make(map[string]string)
	params["id"] = g.Query("id")

	g.JSON(http.StatusOK, test.Get(params))
}

//query 和 post 混合
func QueryPost(g *gin.Context) {
	params := make(map[string]string)
	params["id"] = g.Query("id")
	params["name"] = g.PostForm("name")
	params["age"] = g.PostForm("age")

	g.JSON(http.StatusOK, test.QueryPost(params))
}

//接收 map 的值
func Map(g *gin.Context) {
	name := g.PostForm("name")
	age := g.PostForm("age")
	ids := g.PostFormMap("ids")
	g.JSON(http.StatusOK, test.Map(name, age, ids))

}
