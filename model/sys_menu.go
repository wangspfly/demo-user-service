package model

func init() {
	AutoMigrate = append(AutoMigrate, &SysMenu{})
}

type SysMenu struct {
	BaseModel
	ID       uint      `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name" gorm:"size:128;comment:路由name"`
	Title    string    `json:"title" gorm:"size:128;comment:菜单名称"`
	Icon     string    `json:"icon" gorm:"size:128;comment:菜单图标"`
	Path     string    `json:"path" gorm:"size:128;comment:路由path"`
	Type     string    `json:"type" gorm:"size:1;comment:菜单类型"`
	ParentId uint      `json:"parentId" gorm:"comment:父菜单ID"`
	Sort     int       `json:"sort" gorm:"size:4;"`
	Visible  bool      `json:"visible" gorm:"size:1;"`
	Children []SysMenu `json:"children,omitempty" gorm:"-"`
}
