package router

import (
	"demo-user-service/api"
	"demo-user-service/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routers = append(routers, registerRoleRouter)
}

func registerRoleRouter(router *gin.RouterGroup) {
	r := router.Group("role").Use(middleware.JWT()).Use(middleware.CasbinHandler())
	{
		r.POST("create", api.CreateRole)
		r.PUT("update", api.UpdateRole)
		r.DELETE("delete", api.DeleteRole)
		r.PUT("bindApi", api.UpdatePermissions)
		r.PUT("bindMenu", api.BindMenu)
		r.GET("listApiByRoleKey", api.ListApiByRoleKey)
		r.GET("paging", api.PagingRole)
	}
}
