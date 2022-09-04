package model

func init() {
	AutoMigrate = append(AutoMigrate, &SysApi{})
}

type SysApi struct {
	BaseModel
	ID          uint   `json:"id" gorm:"primaryKey"`
	Path        string `json:"path"`        // api路径
	Description string `json:"description"` // api中文描述
	ApiGroup    string `json:"apiGroup"`    // api组
	Method      string `json:"method"`      // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
}
