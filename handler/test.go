package handler

import (
	"github.com/gin-gonic/gin"
	"goGINLearn/model/test"
	"net/http"
)

func Index(g *gin.Context) {
	g.JSON(http.StatusOK, test.Index())
}
