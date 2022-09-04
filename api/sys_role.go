package api

import (
	"demo-user-service/model"
	"demo-user-service/model/request"
	"demo-user-service/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateRole
// @Tags 角色管理
// @Summary 创建角色
// @Param default body model.SysRole true "角色信息"
// @Success 200 {object} model.SysRole
// @Security Bearer
// @Router /role/create [post]
func CreateRole(c *gin.Context) {
	var role model.SysRole
	err := c.Bind(&role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	create, err := service.CreateRole(role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, create)
}

// PagingRole
// @Tags 角色管理
// @Summary 分页获取角色
// @Param page query int false "页码"
// @Param size query int false "条数"
// @Success 200 {array} model.SysRole
// @Security Bearer
// @Router /role/paging [get]
func PagingRole(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	roles, total, err := service.PagingRole(page, size)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"content": roles,
		"total":   total,
	})
}

// UpdateRole
// @Tags 角色管理
// @Summary 更新角色信息
// @Param default body model.SysRole true "角色信息"
// @Success 200 {object} model.SysRole
// @Security Bearer
// @Router /role/update [put]
func UpdateRole(c *gin.Context) {
	var role model.SysRole
	err := c.Bind(&role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	update, err := service.UpdateRole(role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, update)
}

// DeleteRole
// @Tags 角色管理
// @Summary 删除角色
// @Param default body model.SysRole true "角色信息"
// @Success 200 {string} string
// @Security Bearer
// @Router /role/delete [delete]
func DeleteRole(c *gin.Context) {
	var role model.SysRole
	err := c.Bind(&role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	err = service.DeleteRole(&role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "success")
}

// BindMenu
// @Tags 角色管理
// @Summary 角色绑定菜单权限
// @Param default body request.BindMenuInfo true "角色菜单信息"
// @Success 200 {string} string
// @Security Bearer
// @Router /role/bindMenu [put]
func BindMenu(c *gin.Context) {
	var roleMenu request.BindMenuInfo
	err := c.BindJSON(&roleMenu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	if err := service.BindMenu(roleMenu.RoleKey, roleMenu.Menus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}
