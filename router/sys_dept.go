package router

import (
	"demo-user-service/api"
	"demo-user-service/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routers = append(routers, registerSysDeptRouter)
}

func registerSysDeptRouter(v1 *gin.RouterGroup) {
	r := v1.Group("dept").Use(middleware.JWT()).Use(middleware.CasbinHandler())
	{
		r.POST("create", api.CreateDept)
		r.DELETE("delete", api.DeleteDept)
		r.PUT("update", api.UpdateDept)
		r.GET("tree", api.GetDeptTree)
	}
}
