package router

import (
	"demo-user-service/api"
	"demo-user-service/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routers = append(routers, registerSysUserRouter)
}

func registerSysUserRouter(v1 *gin.RouterGroup) {
	base := v1.Group("user")
	{
		base.POST("login", api.Login)
		base.POST("logout", api.Logout)
		base.POST("refreshToken", api.RefreshToken)
		base.POST("register", api.Register)
		base.POST("resetPassword", api.ResetPassword)
	}
	r := v1.Group("user").Use(middleware.JWT()).Use(middleware.CasbinHandler())
	{
		r.PUT("set", api.SetUser)
		r.GET("paging", api.PagingUser)
		r.DELETE("delete", api.DeleteUser)
	}
}
