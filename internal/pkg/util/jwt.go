package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// 生成JWT
// @param userID 用户ID
// @param secretKey 密钥
// @return token 字符串
func GenerateJWT(userID, secretKey string) (string, error) {
	claims := User{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
		},
	}

	// 使用HS256签名算法
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString([]byte(secretKey))

	return s, err
}

// 解析JWT
// @param jwtString JWT字符串
// @param secretKey 密钥
// @return User用户信息
func ParseJwt(jwtString, secretKey string) (*User, error) {
	t, err := jwt.ParseWithClaims(jwtString, &User{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(*User); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
