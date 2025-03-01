package common

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"iniyou.com/model"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.RegisteredClaims
}

func ReleaseToken(user model.User) (string, error) {
	// 检查用户 ID 是否有效
	if user.ID == 0 {
		return "", fmt.Errorf("invalid user ID")
	}
	// token过期时间
	expiractionTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiractionTime), // 使用 jwt.NewNumericDate
			Issuer:    "oceanlearn.tech",                   // 谁发放的这个Token
			Subject:   "user Token",                        // Token的主题
		},
	}
	// 创建 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名并生成令牌字符串
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
