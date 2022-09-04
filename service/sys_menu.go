package service

import (
	"demo-user-service/global"
	"demo-user-service/model"
	"errors"

	"gorm.io/gorm"
)

func CreateMenu(menu model.SysMenu) (*model.SysMenu, error) {
	if menu.Path == "" || menu.Name == "" || menu.Title == "" {
		return nil, errors.New("menu param is not valid")
	}
	var menuInDB model.SysMenu
	err := global.DB.Where("name = ?", menu.Name).First(&menuInDB).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return &menuInDB, errors.New("存在相同按钮，请修改")
	}
	err = global.DB.Create(&menu).Error
	return &menu, err
}

func DeleteMenu(id uint) error {
	err := global.DB.Where("parent_id = ?", id).First(&model.SysMenu{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("此菜单存在子菜单不可删除")
	}
	var menu model.SysMenu
	return global.DB.Model(&menu).Delete(&menu, id).Error
}

func UpdateMenu(menu model.SysMenu) error {
	if menu.ID == 0 || menu.Path == "" || menu.Name == "" || menu.Title == "" {
		return errors.New("menu param is not valid")
	}
	var oldMenu model.SysMenu
	db := global.DB.Where("id = ?", menu.ID).Find(&oldMenu)
	if oldMenu.Name != menu.Name {
		err := global.DB.Where("id <> ? AND name = ?", menu.ID, menu.Name).First(&model.SysMenu{}).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("存在相同name修改失败")
		}
	}
	return db.Updates(&menu).Error
}

func GetMenu(id uint) (model.SysMenu, error) {
	var menu model.SysMenu
	err := global.DB.Where("id = ?", id).First(&menu).Error
	return menu, err
}

func GetMenuTree() ([]model.SysMenu, error) {
	treeMap, err := getMenusTreeMap()
	if err != nil {
		return nil, err
	}
	menus := treeMap[0]
	for i := 0; i < len(menus); i++ {
		getChildrenList(&menus[i], treeMap)
	}
	if menus == nil {
		menus = []model.SysMenu{}
	}
	return menus, nil
}

func ListMenuByRole(roleKey string) ([]model.SysMenu, error) {
	var role model.SysRole
	err := global.DB.Model(&role).Where("role_key = ?", roleKey).Preload("SysMenus").Find(&role).Error
	if role.SysMenus == nil {
		role.SysMenus = []model.SysMenu{}
	}
	return role.SysMenus, err
}

func GetTreeByRole(roleKey string) ([]model.SysMenu, error) {
	var role model.SysRole
	err := global.DB.Model(&role).Where("role_key = ?", roleKey).Preload("SysMenus").Find(&role).Error
	if err != nil {
		return nil, err
	}
	menusByRole := role.SysMenus
	treeMap := make(map[uint][]model.SysMenu)
	for _, menu := range menusByRole {
		treeMap[menu.ParentId] = append(treeMap[menu.ParentId], menu)
	}
	menus := treeMap[0]
	for i := 0; i < len(menus); i++ {
		getChildrenList(&menus[i], treeMap)
	}
	if menus == nil {
		menus = []model.SysMenu{}
	}
	return menus, nil
}

func getChildrenList(menu *model.SysMenu, treeMap map[uint][]model.SysMenu) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		getChildrenList(&menu.Children[i], treeMap)
	}
}

func getMenusTreeMap() (map[uint][]model.SysMenu, error) {
	var allMenus []model.SysMenu
	treeMap := make(map[uint][]model.SysMenu)
	err := global.DB.Find(&allMenus).Error
	for _, menu := range allMenus {
		treeMap[menu.ParentId] = append(treeMap[menu.ParentId], menu)
	}
	return treeMap, err
}
