/*
@Time : 2024/12/9 15:51
@Author : linx
@File : jwt.go
@dsc: JWT生成和解析
*/

package pkg

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	"time"
)

type MemberClaims struct {
	UserId int64 `json:"-"`
	jwt.StandardClaims
}

type Token struct {
}

func NewToken() *Token {
	return &Token{}
}

const secret = "wjy_blog"

// GenerateToken  生成并返回一个token，以及可能遇到的error
func (t *Token) GenerateToken(memberId int64) (token string, err error) {
	//1.根据当前系统时间，设置生成 token 的到期时间
	now := time.Now()
	expireAt := now.Add((3 * 24) * time.Hour)
	claims := MemberClaims{
		UserId: memberId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: now.Unix(),
			ExpiresAt: expireAt.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    "wjy_blog",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if token, err = tokenClaims.SignedString([]byte(secret)); err != nil {
		zap.L().Error("SignedString Error:", zap.Error(err))
		return "", err
	}

	return token, nil
}

// ParseToken  解析token ，判断token是否合法
func (t *Token) ParseToken(token string) (any, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MemberClaims{}, func(*jwt.Token) (interface{}, error) {
		return secret, nil //使用自定义的秘钥来进行token解密
	})
	if err != nil {
		return nil, err
	}

	claims, ok := tokenClaims.Claims.(*MemberClaims)
	if !ok || !tokenClaims.Valid {
		err = errors.New("用户鉴权失败! ")
		return nil, err
	}

	return claims, err
}
