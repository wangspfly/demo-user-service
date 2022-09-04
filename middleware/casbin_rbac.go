package middleware

import (
	"demo-user-service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		obj := c.Request.URL.Path
		act := c.Request.Method
		e := service.Casbin()
		sub := c.MustGet("username")
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if success {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, gin.H{"msg": "权限不足"})
			c.Abort()
			return
		}
	}
}
