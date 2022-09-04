package service

import (
	"demo-user-service/global"
	"demo-user-service/model"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

func CreateDept(dept model.SysDept) error {
	if dept.Name == "" || (dept.ParentId == 0 && dept.Name != "root") {
		return errors.New("dept param is not valid")
	}
	err := global.DB.Where("name = ?", dept.Name).First(&model.SysDept{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("存在相同部门名称，请修改")
	}
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err := global.DB.Create(&dept).Error
		if err != nil {
			return err
		}
		deptPath := strconv.Itoa(int(dept.ID)) + "/"
		if dept.ParentId != 0 {
			var deptP model.SysDept
			tx.First(&deptP, dept.ParentId)
			deptPath = deptP.Path + deptPath
		} else {
			deptPath = "/0/" + deptPath
		}
		if err := tx.Model(&dept).Update("path", deptPath).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func DeleteDept(id uint) error {
	err := global.DB.Where("parent_id = ?", id).First(&model.SysDept{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("此部门存在子部门不可删除")
	}
	var dept model.SysDept
	return global.DB.Model(&dept).Delete(&dept, id).Error
}

func UpdateDept(dept model.SysDept) error {
	if dept.ID == 0 || dept.Name == "" {
		return errors.New("menu param is not valid")
	}
	var oldDept model.SysDept
	err := global.DB.Where("id = ?", dept.ID).Find(&oldDept).Error
	if err != nil {
		return err
	}
	if oldDept.Name != dept.Name {
		err := global.DB.Where("id <> ? AND name = ?", dept.ID, dept.Name).First(&model.SysDept{}).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("存在相同name修改失败")
		}
	}
	deptPath := strconv.Itoa(int(dept.ID)) + "/"
	if dept.ParentId != 0 {
		var deptP model.SysDept
		err := global.DB.First(&deptP, dept.ParentId).Error
		if err != nil {
			return err
		}
		deptPath = deptP.Path + deptPath
	} else {
		deptPath = "/0/" + deptPath
	}
	dept.Path = deptPath
	return global.DB.Updates(&dept).Error
}

func GetDeptTree() ([]model.SysDept, error) {
	treeMap, err := getDeptTreeMap()
	if err != nil {
		return nil, err
	}
	departments := treeMap[0]
	for i := 0; i < len(departments); i++ {
		getChildrenDept(&departments[i], treeMap)
	}
	if departments == nil {
		departments = []model.SysDept{}
	}
	return departments, nil
}

func getChildrenDept(dept *model.SysDept, treeMap map[uint][]model.SysDept) {
	dept.Children = treeMap[dept.ID]
	for i := 0; i < len(dept.Children); i++ {
		getChildrenDept(&dept.Children[i], treeMap)
	}
}

func getDeptTreeMap() (map[uint][]model.SysDept, error) {
	var allDept []model.SysDept
	treeMap := make(map[uint][]model.SysDept)
	err := global.DB.Find(&allDept).Error
	for _, dept := range allDept {
		treeMap[dept.ParentId] = append(treeMap[dept.ParentId], dept)
	}
	return treeMap, err
}
