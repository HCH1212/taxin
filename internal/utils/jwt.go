package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var accessTokenKey = []byte(os.Getenv("TOKEN_SECRET"))

type Claims struct {
	UserID string
	jwt.RegisteredClaims
}

// GetToken 生成token
func GetToken(userID string) (string, error) {
	// accessToken过期时间一周
	accessTokenTime := time.Now().Add(7 * 24 * time.Hour)

	accessClaims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "my",
			Subject:   "token",
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	accessTokenStr, err := accessToken.SignedString(accessTokenKey)
	if err != nil {
		return "", err
	}
	return accessTokenStr, nil
}

// ParseAccessToken 解析token
func ParseAccessToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return accessTokenKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 有效
	if token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
