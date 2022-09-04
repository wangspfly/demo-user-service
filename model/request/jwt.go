package request

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	DataScope string
	RoleKey   string
	DeptID    uint
	jwt.StandardClaims
}

type DataPermission struct {
	DataScope string
	Username  string
	RoleKey   string
	DeptId    uint
}
