package request

type Register struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
	NickName string
	RoleKey  string `binding:"required"`
	DeptId   uint   `binding:"required"`
}

type Login struct {
	Username string
	Password string
}

type PasswordReset struct {
	Old string
	New string
}
