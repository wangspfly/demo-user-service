package api

import (
	"demo-user-service/model"
	"demo-user-service/model/request"
	"demo-user-service/service"
	"demo-user-service/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Login
// @Tags 用户管理
// @Summary 用户登录
// @Param default body request.Login true "用户名密码"
// @Success 200 {object} response.TokenInfo
// @Failure 500 {string} string
// @Router /user/login [post]
func Login(c *gin.Context) {
	var login request.Login
	err := c.ShouldBind(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	tokenInfo, err := service.Login(login.Username, login.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "登录失败",
		})
		return
	}

	c.JSON(http.StatusOK, tokenInfo)
}

// Logout
// @Tags 用户管理
// @Summary 注销登录
// @Success 200 {string} string
// @Failure 500 {string} string
// @Security Bearer
// @Router /user/logout [post]
func Logout(c *gin.Context) {
	token, err := util.JwtFromHeader(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	err = service.JsonInBlacklist(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "token作废失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "token作废成功",
	})
}

// Register
// @Tags 用户管理
// @Summary 用户注册
// @Param default body request.Register true "用户信息"
// @Success 200 {string} string
// @Security Bearer
// @Router /user/register [post]
func Register(c *gin.Context) {
	var user request.Register
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	err = service.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "注册失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}

func ResetPassword(c *gin.Context) {
	var reset request.PasswordReset
	username := c.MustGet("username").(string)
	err := c.ShouldBind(&reset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	tokenInfo, err := service.ResetPassword(reset, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "重置失败",
		})
		return
	}
	c.JSON(http.StatusOK, tokenInfo)
}

// SetUser
// @Tags 用户管理
// @Summary 设置用户信息
// @Param default body model.SysUser true "用户信息"
// @Success 200 {object} model.SysUser
// @Security Bearer
// @Router /user/set [put]
func SetUser(c *gin.Context) {
	var user model.SysUser
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	p := util.GetPermissionFromContext(c)
	err = service.SetUser(p, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "设置成功"})
}

// PagingUser
// @Tags 用户管理
// @Summary 分页获取用户
// @Param page query int false "页码"
// @Param size query int false "条数"
// @Success 200 {array} model.SysUser
// @Security Bearer
// @Router /user/paging [get]
func PagingUser(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	deptId, _ := strconv.Atoi(c.Query("deptId"))
	roles, total, err := service.PagingUser(uint(deptId), page, size)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"content": roles,
		"total":   total,
	})
}

// RefreshToken
// @Tags 用户管理
// @Summary 刷新token
// @Success 200 {string} string
// @Security Bearer
// @Router /user/refreshToken [post]
func RefreshToken(c *gin.Context) {
	token, err := util.JwtFromHeader(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	refreshToken, err := service.RefreshToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, refreshToken)
}

// DeleteUser
// @Tags 用户管理
// @Summary 删除用户
// @Param default body request.GetById true "用户ID"
// @Success 200 {string} string
// @Security Bearer
// @Router /user/delete [delete]
func DeleteUser(c *gin.Context) {
	var reqID request.GetById
	err := c.Bind(&reqID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	p := util.GetPermissionFromContext(c)

	err = service.DeleteUser(p, reqID.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
