package config

import "time"

var JWTConfig = struct {
	SigningKey  []byte
	ExpiresTime time.Duration
	Issuer      string
}{
	SigningKey:  []byte("user_center_jwt_secret"),
	ExpiresTime: time.Hour * 3,
	Issuer:      "demo-user-service",
}
