/*
@author:Deng.l.w
@version:1.20
@date:2023-03-06 16:21
@file:jwt.go
*/

package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenExpireDuration = time.Hour * 24 * 365

// mySecret 签名
var mySecret = []byte("dengliwei")

// MyClaims 自定义声明
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims //标准声明
}

// GenToken 生成token
func GenToken(userID int64, username string) (string, error) {
	// 创建自己的声明
	c := MyClaims{
		userID,
		"username", //自定义字段
		jwt.StandardClaims{
			//过期时间
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			// 签发人
			Issuer: "bluebell",
		},
	}

	// 使用指定签名方法创造签名对象 Token{Raw,Method,Header,Claims,Signature,Valid}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定secret签名并获得完整编码后的token
	return token.SignedString(mySecret)
}

// ParseToekn 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid Token")
}
