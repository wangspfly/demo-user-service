package util

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	generateToken, err := GenerateToken("admin", "admin", 111, "")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(generateToken)
}
