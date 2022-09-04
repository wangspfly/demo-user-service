package api

import (
	"demo-user-service/model/request"
	"demo-user-service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdatePermissions
// @Tags 角色管理
// @Summary 角色绑定Api权限
// @Param default body request.CasbinInReceive true "角色API信息"
// @Success 200 {string} string
// @Security Bearer
// @Router /role/bindApi [put]
func UpdatePermissions(c *gin.Context) {
	var cmr request.CasbinInReceive
	err := c.BindJSON(&cmr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	if err := service.UpdatePermissions(cmr.RoleKey, cmr.CasbinInfos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}

// ListApiByRoleKey
// @Tags 角色管理
// @Summary 获取角色绑定的API权限
// @Param roleKey query string true "角色ID"
// @Success 200 {array} request.CasbinInfo
// @Security Bearer
// @Router /role/listApiByRoleKey [get]
func ListApiByRoleKey(c *gin.Context) {
	roleKey, _ := c.GetQuery("roleKey")
	paths := service.ListApiByRoleKey(roleKey)
	c.JSON(http.StatusOK, paths)
}
