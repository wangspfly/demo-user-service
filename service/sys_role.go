package service

import (
	"demo-user-service/global"
	"demo-user-service/model"
	"errors"

	"gorm.io/gorm"
)

func CreateRole(role model.SysRole) (model.SysRole, error) {
	if role.RoleKey == "" || role.RoleName == "" {
		return model.SysRole{}, errors.New("角色参数错误")
	}
	err := global.DB.Where("role_key = ?", role.RoleKey).First(&model.SysRole{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return role, errors.New("存在相同角色")
	}
	err = global.DB.Create(&role).Error
	return role, err
}

func PagingRole(page, size int) ([]model.SysRole, int64, error) {
	var roles []model.SysRole
	var total int64
	tx := global.DB.Where("id <> ?", 1)
	if page > 0 && size > 0 {
		tx = tx.Limit(size).Offset((page - 1) * size)
	} else {
		tx = tx.Limit(10).Offset(0)
	}
	err := tx.Find(&roles).Count(&total).Error
	return roles, total, err
}

func UpdateRole(role model.SysRole) (model.SysRole, error) {
	if role.RoleKey == "" || role.RoleName == "" {
		return model.SysRole{}, errors.New("角色参数错误")
	}
	err := global.DB.Where("role_key = ?", role.RoleKey).First(&model.SysRole{}).Updates(&role).Error
	return role, err
}

func DeleteRole(role *model.SysRole) error {
	err := global.DB.Where("role_key = ?", role.RoleKey).First(&model.SysUser{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	err = global.DB.Unscoped().Delete(role).Error
	if err != nil {
		return err
	}
	DeleteCasbinRole(role.RoleKey)
	return nil
}

func BindMenu(roleKey string, menus []model.SysMenu) error {
	if roleKey == "" {
		return errors.New("参数错误")
	}
	for _, menu := range menus {
		if menu.ID == 0 {
			return errors.New("菜单ID错误")
		}
	}
	var role model.SysRole
	global.DB.Preload("SysMenus").First(&role, "role_key = ?", roleKey)
	return global.DB.Model(&role).Association("SysMenus").Replace(&menus)
}
