package service

import (
	"demo-user-service/global"
	"demo-user-service/model"
	"errors"

	"gorm.io/gorm"
)

func JsonInBlacklist(token string) error {
	if token == "" {
		return errors.New("非法参数")
	}
	jwt := model.JwtBlacklist{Jwt: token}
	err := global.DB.Create(&jwt).Error
	if err != nil {
		return err
	}
	return nil
}

func IsBlacklist(token string) bool {
	err := global.DB.Where("jwt = ?", token).First(&model.JwtBlacklist{}).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}
