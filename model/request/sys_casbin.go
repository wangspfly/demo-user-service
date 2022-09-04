package request

type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

type CasbinInReceive struct {
	RoleKey     string       `json:"roleKey"` // 角色编码
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}
