package model

func init() {
	AutoMigrate = append(AutoMigrate, &SysUser{})
}

type SysUser struct {
	BaseModel
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"not null;unique;size:64;comment:用户名"`
	Password string `json:"-" gorm:"size:128;comment:密码"`
	NickName string `json:"nickName" gorm:"size:128;comment:昵称"`
	Phone    string `json:"phone" gorm:"size:11;comment:手机号"`
	RoleKey  string `json:"roleKey" gorm:"size:90;comment:角色编码"`
	Avatar   string `json:"avatar" gorm:"size:255;comment:头像"`
	Sex      string `json:"sex" gorm:"size:255;comment:性别"`
	Email    string `json:"email" gorm:"size:128;comment:邮箱"`
	DeptId   uint   `json:"deptId" gorm:"size:20;comment:部门"`
	Status   string `json:"status" gorm:"size:4;comment:状态"`
}
