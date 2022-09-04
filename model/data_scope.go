package model

import "time"

func init() {
	AutoMigrate = append(AutoMigrate, &DataScope{})
}

type DataScope struct {
	BaseModel
	Subject   string    `gorm:"comment:数据所有者"`
	Object    string    `gorm:"comment:数据资源"`
	ExpiresAt time.Time `json:"expiresAt" gorm:"comment:过期时间"`
}
