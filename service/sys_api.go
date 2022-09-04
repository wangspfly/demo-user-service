package service

import (
	"demo-user-service/global"
	"demo-user-service/model"
	"strings"

	"github.com/gin-gonic/gin"
)

func RegisterApi(info gin.RoutesInfo) []model.SysApi {
	var apis []model.SysApi
	for _, routeInfo := range info {
		group := getApiGroup(routeInfo.Path)
		if group == "" {
			continue
		}
		api := model.SysApi{
			Path:        routeInfo.Path,
			Description: "",
			ApiGroup:    group,
			Method:      routeInfo.Method,
		}
		apis = append(apis, api)
	}
	global.DB.Exec("truncate table sys_apis")
	global.DB.Create(&apis)
	return apis
}

func ListApi() []model.SysApi {
	var apis []model.SysApi
	global.DB.Find(&apis)
	return apis
}

func getApiGroup(path string) string {
	split := strings.Split(path, "v1")
	if len(split) < 2 {
		return ""
	}
	i := strings.Split(split[1], "/")
	return i[1]
}
