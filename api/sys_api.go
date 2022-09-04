package api

import (
	"demo-user-service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListApi
// @Tags 获取Api列表
// @Summary 获取所有Api
// @Security Bearer
// @Success 200 {array} model.SysApi
// @Router /api/list [get]
func ListApi(c *gin.Context) {
	c.JSON(http.StatusOK, service.ListApi())
}
