package model

func init() {
	AutoMigrate = append(AutoMigrate, &SysDept{})
}

type SysDept struct {
	BaseModel
	ID       uint      `json:"id" gorm:"primaryKey"`
	ParentId uint      `json:"parentId" gorm:""`        //上级部门
	Name     string    `json:"name"  gorm:"size:128;"`  //部门名称
	Path     string    `json:"path" gorm:"size:255;"`   //
	Leader   string    `json:"leader" gorm:"size:128;"` //负责人
	Phone    string    `json:"phone" gorm:"size:11;"`   //手机
	Email    string    `json:"email" gorm:"size:64;"`   //邮箱
	Children []SysDept `json:"children" gorm:"-"`
}
