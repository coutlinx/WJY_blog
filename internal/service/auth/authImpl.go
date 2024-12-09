/*
@Time : 2024/12/9 15:41
@Author : linx
@File : authImpl.go
@dsc: 注册
*/

package auth

import (
	"blog/internal/interface/auth"
	"blog/internal/models"
	"blog/pkg"
	"errors"
)

type Type string

const (
	Email   Type = "email"
	Account Type = "account"
)

func (t Type) GetImpl(user auth.User) auth.Auth {
	switch t {
	case Email:
		return &WithEmail{
			User: user,
		}
	case Account:
		return &WithAccount{
			User: user,
		}
	default:
		return nil
	}
}

type WithEmail struct {
	User auth.User
}

func (r *WithEmail) UserRegister() (string, error) {
	var err error
	// 加密密码
	if err = r.User.EncryptPassword(); err != nil {
		return "", err
	}
	// 创建用户
	var data any
	if data, err = r.User.CreateUser(); err != nil {
		return "", err
	}
	user := data.(*models.BlogUser)

	// 生成Token
	var token string
	if token, err = pkg.NewToken().GenerateToken(user.Id); err != nil {
		return "", err
	}
	return token, nil
}

func (r *WithEmail) UserLogin() (string, error) {
	var err error
	originPassword := r.User.GetPassword()
	var user *models.BlogUser
	var data any
	if data, err = r.User.GetUser(); err != nil {
		return "", err
	}
	user = data.(*models.BlogUser)
	if user.Password != pkg.NewEncrypt().EncryptWithSalt(originPassword, user.Salt) {
		return "", errors.New("账号或密码错误")
	}
	// 生成Token
	var token string
	if token, err = pkg.NewToken().GenerateToken(user.Id); err != nil {
		return "", err
	}
	return token, nil

}

type WithAccount struct {
	User auth.User
}

func (r *WithAccount) UserRegister() (string, error) {
	var err error
	// 加密密码
	if err = r.User.EncryptPassword(); err != nil {
		return "", err
	}
	// 创建用户
	var data any
	if data, err = r.User.CreateUser(); err != nil {
		return "", err
	}
	user := data.(*models.BlogUser)

	// 生成Token
	var token string
	if token, err = pkg.NewToken().GenerateToken(user.Id); err != nil {
		return "", err
	}
	return token, nil
}

func (r *WithAccount) UserLogin() (string, error) {
	var err error
	originPassword := r.User.GetPassword()
	var user *models.BlogUser
	var data any
	if data, err = r.User.GetUser(); err != nil {
		return "", err
	}
	user = data.(*models.BlogUser)
	if user.Password != pkg.NewEncrypt().EncryptWithSalt(originPassword, user.Salt) {
		return "", errors.New("账号或密码错误")
	}
	// 生成Token
	var token string
	if token, err = pkg.NewToken().GenerateToken(user.Id); err != nil {
		return "", err
	}
	return token, nil

}
