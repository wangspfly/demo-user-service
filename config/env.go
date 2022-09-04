package config

import (
	"demo-user-service/global"
)

func init() {
	global.DBAddress = "host=localhost user=postgres password=postgrespw dbname=user_center port=49153 sslmode=disable TimeZone=Asia/Shanghai"
}
