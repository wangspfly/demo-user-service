package config

import (
	"demo-user-service/model"
	"demo-user-service/model/enum"
	"demo-user-service/model/request"
)

const (
	RoleRoot     = "root"
	RoleAdmin    = "admin"
	RoleReviewer = "reviewer"
	RoleCommon   = "common"
)

var Dept = model.SysDept{
	ID:   1,
	Name: "root",
	Path: "/0/",
}

var Roles = []model.SysRole{
	{
		ID:        1,
		RoleKey:   RoleRoot,
		RoleName:  "超级管理员",
		DataScope: enum.AllDept,
	}, {
		ID:        2,
		RoleKey:   RoleAdmin,
		RoleName:  "管理员",
		DataScope: enum.DeptTree,
	}, {
		ID:       3,
		RoleKey:  RoleReviewer,
		RoleName: "审核员",
	}, {
		ID:       4,
		RoleKey:  RoleCommon,
		RoleName: "技术员",
	},
}

var User = request.Register{
	Username: "admin",
	Password: "admin",
	NickName: "管理员",
	RoleKey:  RoleRoot,
	DeptId:   0,
}

var Menus = []model.SysMenu{
	{
		ID:       1,
		Name:     "operation",
		Title:    "实时数据",
		Path:     "operation",
		Type:     "",
		ParentId: 0,
		Sort:     1,
		Visible:  true,
	}, {
		ID:       2,
		Name:     "classify",
		Title:    "基础数据",
		Path:     "classify/list",
		Type:     "",
		ParentId: 0,
		Sort:     2,
		Visible:  true,
	}, {
		ID:       3,
		Name:     "model",
		Title:    "模型管理",
		Path:     "model",
		Type:     "",
		ParentId: 0,
		Sort:     3,
		Visible:  true,
	}, {
		ID:       4,
		Name:     "approval",
		Title:    "审批中心",
		Path:     "approval/list",
		Type:     "",
		ParentId: 0,
		Sort:     4,
		Visible:  true,
	}, {
		ID:       5,
		Name:     "userEmpty",
		Title:    "权限管理",
		Path:     "user",
		Type:     "",
		ParentId: 0,
		Sort:     5,
		Visible:  true,
		Children: []model.SysMenu{
			{
				ID:       6,
				Name:     "role",
				Title:    "角色管理",
				Path:     "user/role",
				Type:     "",
				ParentId: 5,
				Sort:     1,
				Visible:  true,
			}, {
				ID:       7,
				Name:     "user",
				Title:    "用户管理",
				Path:     "user/user",
				Type:     "",
				ParentId: 5,
				Sort:     2,
				Visible:  true,
			}, {
				ID:       8,
				Name:     "menu",
				Title:    "菜单管理",
				Path:     "user/menu",
				Type:     "",
				ParentId: 5,
				Sort:     3,
				Visible:  true,
			},
		},
	}, {
		ID:       9,
		Name:     "message",
		Title:    "消息中心",
		Path:     "message",
		Type:     "",
		ParentId: 0,
		Sort:     6,
		Visible:  true,
	}, {
		ID:       10,
		Name:     "dept",
		Title:    "组织架构",
		Path:     "dept",
		Type:     "",
		ParentId: 0,
		Sort:     7,
		Visible:  true,
	},
}
