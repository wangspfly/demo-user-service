package model

func init() {
	AutoMigrate = append(AutoMigrate, &SysRole{})
}

type SysRole struct {
	BaseModel
	ID        uint      `json:"id" gorm:"primaryKey"`
	RoleKey   string    `json:"roleKey" gorm:"not null;unique;comment:角色代码;size:90"`
	RoleName  string    `json:"roleName" gorm:"size:128;comment:角色名称"`
	Status    string    `json:"status" gorm:"size:4;"`
	Remark    string    `json:"remark" gorm:"size:255;"`
	DataScope string    `json:"dataScope" gorm:"comment:数据权限"`
	SysMenus  []SysMenu `json:"menus" gorm:"many2many:sys_role_menus;foreignKey:RoleKey"`
}
