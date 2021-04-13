package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	//test say hello
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin")
	})

	router.Run()
}
