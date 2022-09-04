package service

import (
	"demo-user-service/config"
	"demo-user-service/global"
	"demo-user-service/model/request"
	"errors"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDBWithCustomTable(global.DB, &gormadapter.CasbinRule{})
		m, err := model.NewModelFromString(config.Rbac)
		if err != nil {
			panic(err)
		}
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(m, a)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}

func UpdatePermissions(roleKey string, infos []request.CasbinInfo) error {
	if roleKey == "" {
		return errors.New("参数错误")
	}
	_, _ = Casbin().DeletePermissionsForUser("role:" + roleKey)
	var permissions [][]string
	for _, info := range infos {
		if info.Path == "" || info.Method == "" {
			continue
		}
		permissions = append(permissions, []string{info.Path, info.Method})
	}
	success, err := Casbin().AddPermissionsForUser("role:"+roleKey, permissions...)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

func ListApiByRoleKey(roleKey string) []request.CasbinInfo {
	if roleKey == "" {
		return nil
	}
	paths := make([]request.CasbinInfo, 0)
	policy := GetPermissionsForRole(roleKey)
	for _, v := range policy {
		paths = append(paths, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return paths
}

func GetPermissionsForRole(roleKey string) [][]string {
	return Casbin().GetPermissionsForUser("role:" + roleKey)
}

func AddRoleForUser(username, roleKey string) {
	_, _ = Casbin().AddRoleForUser("user:"+username, "role:"+roleKey)
}

func DeleteRolesForUser(username string) {
	_, _ = Casbin().DeleteRolesForUser("user:" + username)
}

func DeleteCasbinUser(username string) {
	_, _ = Casbin().DeleteUser("user:" + username)
}

func DeleteCasbinRole(roleKey string) {
	_, _ = Casbin().DeleteRole("role:" + roleKey)
}
