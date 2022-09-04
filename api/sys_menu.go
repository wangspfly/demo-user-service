package api

import (
	"demo-user-service/model"
	"demo-user-service/model/request"
	"demo-user-service/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateMenu
// @Tags 菜单管理
// @Summary 创建菜单
// @Param default body model.SysMenu true "菜单信息"
// @Success 200 {object} model.SysMenu
// @Security Bearer
// @Router /menu/create [post]
func CreateMenu(c *gin.Context) {
	var menu model.SysMenu
	err := c.Bind(&menu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	create, err := service.CreateMenu(menu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, create)
}

// DeleteMenu
// @Tags 菜单管理
// @Summary 删除菜单
// @Param default body request.GetById true "菜单ID"
// @Success 200 {string} string
// @Security Bearer
// @Router /menu/delete [delete]
func DeleteMenu(c *gin.Context) {
	var menu request.GetById
	err := c.Bind(&menu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	err = service.DeleteMenu(menu.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}

// UpdateMenu
// @Tags 菜单管理
// @Summary 更新菜单
// @Param default body model.SysMenu true "菜单信息"
// @Success 200 {string} string
// @Security Bearer
// @Router /menu/update [put]
func UpdateMenu(c *gin.Context) {
	var menu model.SysMenu
	err := c.Bind(&menu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	err = service.UpdateMenu(menu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "更新失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}

// GetMenuTree
// @Tags 菜单管理
// @Summary 获取菜单树
// @Success 200 {array} model.SysMenu
// @Security Bearer
// @Router /menu/getTree [get]
func GetMenuTree(c *gin.Context) {
	menus, err := service.GetMenuTree()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, menus)
}

// GetMenu
// @Tags 菜单管理
// @Summary 获取菜单详情
// @Param id query integer true "菜单编号"
// @Success 200 {object} model.SysMenu
// @Security Bearer
// @Router /menu/get [get]
func GetMenu(c *gin.Context) {
	queryID := c.Query("id")
	id, err := strconv.Atoi(queryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	menu, err := service.GetMenu(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, menu)
}

// ListMenuByRole
// @Tags 菜单管理
// @Summary 获取角色菜单树
// @Param roleKey query string true "角色"
// @Success 200 {array} model.SysMenu
// @Security Bearer
// @Router /menu/listByRole [get]
func ListMenuByRole(c *gin.Context) {
	roleKey := c.Query("roleKey")
	menus, err := service.ListMenuByRole(roleKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, menus)
}

// GetTreeByRole
// @Tags 菜单管理
// @Summary 获取角色菜单树
// @Success 200 {array} model.SysMenu
// @Security Bearer
// @Router /menu/getTreeByRole [get]
func GetTreeByRole(c *gin.Context) {
	roleKey := c.MustGet("roleKey").(string)
	menus, err := service.GetTreeByRole(roleKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, menus)
}
