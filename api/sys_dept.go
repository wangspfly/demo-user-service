package api

import (
	"demo-user-service/model"
	"demo-user-service/model/request"
	"demo-user-service/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateDept
// @Tags 部门管理
// @Summary 添加部门
// @Param default body model.SysDept true "部门信息"
// @Success 200 {object} model.SysDept
// @Security Bearer
// @Router /dept/create [post]
func CreateDept(c *gin.Context) {
	var dept model.SysDept
	err := c.ShouldBind(&dept)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	err = service.CreateDept(dept)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "创建失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "创建成功",
	})
}

// DeleteDept
// @Tags 部门管理
// @Summary 添加部门
// @Param default body request.GetById true "部门ID"
// @Success 200 {string} string
// @Security Bearer
// @Router /dept/delete [delete]
func DeleteDept(c *gin.Context) {
	var byID request.GetById
	err := c.Bind(&byID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	err = service.DeleteDept(byID.ID)
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

// UpdateDept
// @Tags 部门管理
// @Summary 更新部门
// @Param default body model.SysDept true "部门信息"
// @Success 200 {string} string
// @Security Bearer
// @Router /dept/update [put]
func UpdateDept(c *gin.Context) {
	var dept model.SysDept
	err := c.Bind(&dept)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	err = service.UpdateDept(dept)
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

// GetDeptTree
// @Tags 部门管理
// @Summary 获取部门树
// @Success 200 {array} model.SysDept
// @Security Bearer
// @Router /dept/tree [get]
func GetDeptTree(c *gin.Context) {
	tree, err := service.GetDeptTree()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tree)
}
