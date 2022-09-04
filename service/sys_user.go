package service

import (
	"demo-user-service/config"
	"demo-user-service/global"
	"demo-user-service/model"
	"demo-user-service/model/request"
	"demo-user-service/model/response"
	"demo-user-service/util"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

// Login checks if authentication information exists
func Login(username, password string) (response.TokenInfo, error) {
	var user model.SysUser
	password = util.EncodeMD5(password)
	err := global.DB.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return response.TokenInfo{}, err
	}

	if user.ID <= 0 {
		return response.TokenInfo{}, errors.New("user not found")
	}
	var role model.SysRole
	err = global.DB.Where("role_key = ?", user.RoleKey).Find(&role).Error
	if err != nil {
		return response.TokenInfo{}, err
	}
	token, err := util.GenerateToken(user.Username, user.RoleKey, user.DeptId, role.DataScope)
	if err != nil {
		return token, err
	}
	return token, err
}

func Register(r request.Register) error {
	if !errors.Is(global.DB.Where("username = ?", r.Username).First(&model.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已注册")
	}
	user := &model.SysUser{
		Username: r.Username,
		Password: util.EncodeMD5(r.Password),
		RoleKey:  r.RoleKey,
		DeptId:   r.DeptId,
	}
	err := global.DB.Create(&user).Error
	if err != nil {
		return err
	}
	AddRoleForUser(user.Username, user.RoleKey)
	return nil
}

func ResetPassword(r request.PasswordReset, username string) (response.TokenInfo, error) {
	var user model.SysUser
	old := util.EncodeMD5(r.Old)
	err := global.DB.Where("username = ? AND password = ?", username, old).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return response.TokenInfo{}, err
	}

	if user.ID <= 0 {
		return response.TokenInfo{}, errors.New("user not found")
	}
	user.Password = util.EncodeMD5(r.New)
	global.DB.Save(&user)
	var role model.SysRole
	err = global.DB.Where("role_key = ?", user.RoleKey).Find(&role).Error
	if err != nil {
		return response.TokenInfo{}, err
	}
	token, err := util.GenerateToken(user.Username, user.RoleKey, user.DeptId, role.DataScope)
	if err != nil {
		return token, err
	}
	return token, err
}

func PagingUser(deptId uint, page, size int) ([]model.SysUser, int64, error) {
	var users []model.SysUser
	var total int64
	tx := global.DB
	if deptId != 0 {
		tx = tx.Where("sys_users.dept_id in(select id from sys_dept where path like ? )",
			"%/"+strconv.Itoa(int(deptId))+"/%")
	}
	if page > 0 && size > 0 {
		tx = tx.Limit(size).Offset((page - 1) * size)
	} else {
		tx = tx.Limit(10).Offset(0)
	}
	err := tx.Find(&users).Count(&total).Error
	return users, total, err
}

func RefreshToken(token string) (response.TokenInfo, error) {
	claims, err := util.ParseToken(token)
	if err != nil {
		return response.TokenInfo{}, err
	}
	newToken, err := util.GenerateToken(claims.Subject, claims.RoleKey, claims.DeptID, claims.DataScope)
	if err != nil {
		return response.TokenInfo{}, err
	}
	return newToken, nil
}

func SetUser(p *request.DataPermission, user model.SysUser) error {
	if user.ID == 0 {
		return errors.New("user id is null")
	}
	if user.Username == "" {
		return errors.New("username is null")
	}
	if user.RoleKey == "" || user.RoleKey == config.RoleRoot {
		return errors.New("no permission to set this role")
	}

	var oldUser model.SysUser
	first := global.DB.Scopes(util.UserPermission(p)).First(&oldUser, user.ID)
	if first.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}

	updates := global.DB.Updates(&user)
	if err := updates.Error; err != nil {
		return err
	}
	if oldUser.RoleKey == user.RoleKey {
		return nil
	}
	DeleteRolesForUser(user.Username)
	AddRoleForUser(user.Username, user.RoleKey)
	return nil
}

func DeleteUser(p *request.DataPermission, userID uint) error {
	var user model.SysUser
	tx := global.DB.Scopes(util.UserPermission(p)).Where("id = ?", userID).Delete(&user)
	if err := tx.Error; err != nil {
		return err
	}
	if tx.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	DeleteCasbinUser(user.Username)
	return nil
}
