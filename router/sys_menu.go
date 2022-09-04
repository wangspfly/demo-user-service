package router

import (
	"demo-user-service/api"
	"demo-user-service/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routers = append(routers, registerMenuRouter)
}

func registerMenuRouter(router *gin.RouterGroup) {
	public := router.Group("menu").Use(middleware.JWT())
	{
		public.GET("getTree", api.GetMenuTree)
		public.GET("listByRole", api.ListMenuByRole)
		public.GET("getTreeByRole", api.GetTreeByRole)
	}
	r := router.Group("menu").Use(middleware.JWT()).Use(middleware.CasbinHandler())
	{
		r.POST("create", api.CreateMenu)
		r.DELETE("delete", api.DeleteMenu)
		r.PUT("update", api.UpdateMenu)
		r.GET("get", api.GetMenu)
	}
}
