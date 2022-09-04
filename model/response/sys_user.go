package response

import "time"

type TokenInfo struct {
	Token  string    `json:"token"`  // 访问令牌
	Expire time.Time `json:"expire"` // 令牌到期时间
}
