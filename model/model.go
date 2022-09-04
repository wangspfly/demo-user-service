package model

import (
	"database/sql/driver"
	"demo-user-service/global"
	"errors"
	"fmt"
	"strings"
	"time"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var AutoMigrate = []interface{}{&gormadapter.CasbinRule{}}

func InitDB() {
	var err error
	global.DB, err = gorm.Open(postgres.Open(global.DBAddress), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)},
	)
	if err != nil {
		fmt.Println("数据库连接出现了问题: ", err)
		return
	}
	err = global.DB.AutoMigrate(AutoMigrate...)
	if err != nil {
		fmt.Println("更新表结构失败: ", err)
		return
	}
}

type BaseModel struct {
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:最后更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}

const TimeFormat = "2006-01-02 15:04:05"

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	jsonTime, err := time.Parse(TimeFormat, timeStr)
	*t = Timestamp(jsonTime)
	return err
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format(TimeFormat))
	return []byte(formatted), nil
}

func (t Timestamp) Value() (driver.Value, error) {
	return time.Time(t).Format(TimeFormat), nil
}

func (t *Timestamp) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		*t = Timestamp(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t Timestamp) String() string {
	return time.Time(t).Format(TimeFormat)
}
