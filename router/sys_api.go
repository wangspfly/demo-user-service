package router

import (
	"demo-user-service/api"
	"demo-user-service/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routers = append(routers, registerApiRouter)
}

func registerApiRouter(router *gin.RouterGroup) {
	r := router.Group("api").Use(middleware.JWT()).Use(middleware.Cors())
	{
		r.GET("list", api.ListApi)
	}
}
