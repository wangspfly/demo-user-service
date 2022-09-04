package middleware

import (
	"demo-user-service/service"
	"demo-user-service/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := util.JwtFromHeader(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": err.Error(),
			})
			c.Abort()
			return
		}
		if service.IsBlacklist(token) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "您的帐户异地登陆或令牌失效",
			})
			c.Abort()
			return
		}

		claims, err := util.ParseToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("username", claims.Subject)
		c.Set("roleKey", claims.RoleKey)
		c.Set("deptID", claims.DeptID)
		c.Set("dataScope", claims.DataScope)
		c.Next()
	}
}
