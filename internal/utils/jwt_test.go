package utils

import (
	"testing"
)

func TestJWT(t *testing.T) {
	// 生成token
	userID := "123456"
	token, err := GetToken(userID)
	if err != nil {
		t.Error(err)
	}
	t.Log(token)

	// 解析token
	claims, err := ParseAccessToken(token)
	if err != nil {
		t.Error(err)
	}
	t.Log(claims.UserID)
}
