package request

import "demo-user-service/model"

type BindMenuInfo struct {
	RoleKey string          `json:"roleKey"`
	Menus   []model.SysMenu `json:"menus"`
}
