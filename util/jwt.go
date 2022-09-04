package util

import (
	"demo-user-service/config"
	"demo-user-service/model/enum"
	"demo-user-service/model/request"
	"demo-user-service/model/response"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GenerateToken(username, roleKey string, deptID uint, dataScope string) (response.TokenInfo, error) {
	jwtConfig := config.JWTConfig
	now := time.Now()
	expire := now.Add(jwtConfig.ExpiresTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &request.Claims{
		DataScope: dataScope,
		RoleKey:   roleKey,
		DeptID:    deptID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    jwtConfig.Issuer,
			NotBefore: now.Unix() - 1000,
			Subject:   "user:" + username,
		},
	})
	tokenString, err := token.SignedString(config.JWTConfig.SigningKey)
	if err != nil {
		return response.TokenInfo{}, err
	}
	tokenInfo := response.TokenInfo{
		Token:  tokenString,
		Expire: expire,
	}
	return tokenInfo, nil
}

// ParseToken parsing token
func ParseToken(tokenString string) (*request.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return config.JWTConfig.SigningKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token.Claims.(*request.Claims), nil
}

func JwtFromHeader(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("auth header is empty")
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", errors.New("auth header is invalid")
	}

	return parts[1], nil
}

func GetPermissionFromContext(c *gin.Context) *request.DataPermission {
	return &request.DataPermission{
		DataScope: c.MustGet("dataScope").(string),
		Username:  c.MustGet("username").(string),
		RoleKey:   c.MustGet("roleKey").(string),
		DeptId:    c.MustGet("deptID").(uint),
	}
}

func UserPermission(p *request.DataPermission) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch p.DataScope {
		case enum.DeptTree:
			return db.Where("sys_users.dept_id in(select id from sys_depts where path like ? )", "%/"+strconv.Itoa(int(p.DeptId))+"/%")
		case enum.AllDept:
			return db
		default:
			return db.Where("sys_users.id = ", 0)
		}
	}
}
