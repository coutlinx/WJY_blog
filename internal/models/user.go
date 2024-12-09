/*
@Time : 2024/12/6 14:42
@Author : linx
@File : user.go
@dsc: 用户
*/

package models

import (
	"blog/configs"
	"blog/pkg"
)

type BlogUser struct {
	BaseModel
	UserName string `json:"user_name" gorm:"column:user_name;"`
	Account  string `json:"account" gorm:"column:account;"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
	Salt     string `json:"salt" gorm:"column:salt"`
}

func (receiver *BlogUser) TableName() string {
	return "blog_user"
}

func (receiver *BlogUser) GetPassword() string {
	return receiver.Password
}

func (receiver *BlogUser) EncryptPassword() error {
	// 生成Salt
	receiver.Salt = pkg.NewRandomGenerate().GenerateSalt(16)
	receiver.Password = pkg.NewEncrypt().EncryptWithSalt(receiver.Password, receiver.Salt)
	return nil
}

func (receiver *BlogUser) CreateUser() (any, error) {
	if err := configs.Config.DB.Model(&BlogUser{}).Create(receiver).Error; err != nil {
		return nil, err
	}
	return receiver, nil
}

func (receiver *BlogUser) GetUser() (any, error) {
	if err := configs.Config.DB.Model(&BlogUser{}).Where("account = ? OR email = ?", receiver.UserName, receiver.Email).First(receiver).Error; err != nil {
		return nil, err
	}
	return receiver, nil
}
