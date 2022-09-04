package service

import (
	"demo-user-service/config"
	"demo-user-service/global"
	"demo-user-service/model"
	"demo-user-service/model/request"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func InitData(info gin.RoutesInfo) {
	// 初始化API
	apis := RegisterApi(info)

	// 创建绑定按钮
	menus := getMenus(config.Menus)
	err := global.DB.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(menus, len(menus)).Error
	if err != nil {
		log.Println(err.Error())
	}

	// 创建角色
	err2 := global.DB.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(config.Roles, 10).Error
	if err2 != nil {
		log.Println(err2.Error())
	}
	_ = BindMenu(config.RoleRoot, menus)
	_ = BindMenu(config.RoleAdmin, menus)
	var infos []request.CasbinInfo
	for _, api := range apis {
		infos = append(infos, request.CasbinInfo{
			Path:   api.Path,
			Method: api.Method,
		})
	}
	_ = UpdatePermissions(config.RoleAdmin, infos)

	_ = CreateDept(config.Dept)

	// 创建用户
	err3 := Register(config.User)
	if err3 != nil {
		log.Println(err3.Error())
	}
}

func getMenus(data []model.SysMenu) []model.SysMenu {
	var menus []model.SysMenu
	for _, menu := range data {
		menus = append(menus, menu)
		if menu.Children != nil {
			menus = append(menus, getMenus(menu.Children)...)
		}
	}
	return menus
}
