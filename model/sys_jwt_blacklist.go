package model

func init() {
	AutoMigrate = append(AutoMigrate, &JwtBlacklist{})
}

type JwtBlacklist struct {
	BaseModel
	ID  uint   `gorm:"primaryKey"`
	Jwt string `gorm:"type:text;comment:jwt"`
}
