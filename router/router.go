package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	routers = make([]func(*gin.RouterGroup), 0)
)

var r = gin.Default()

func Router() *gin.Engine {
	r.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, "ok")
	})

	v1 := r.Group("/api/v1")
	for _, f := range routers {
		f(v1)
	}
	return r
}
