/*
@Time : 2024/12/9 15:41
@Author : linx
@File : service.go
@dsc:
*/

package auth

import (
	"blog/internal/models"
	"blog/internal/models/request"
	"blog/internal/models/response"
	"github.com/gin-gonic/gin"
)

type Service struct {
}

func NewAuthService() *Service {
	return &Service{}
}

// RegisterWithAccountHandle 邮箱注册
func (receiver *Service) RegisterWithAccountHandle(c *gin.Context, req request.RegisterReq) (resp *response.RegisterResp, err error) {
	resp = new(response.RegisterResp)
	var user = &models.BlogUser{
		Account:  req.Account,
		Password: req.Password,
	}
	if resp.Token, err = Account.GetImpl(user).UserRegister(); err != nil {
		return nil, err
	}
	return
}

// RegisterWithEmailHandle 账号注册
func (receiver *Service) RegisterWithEmailHandle(c *gin.Context, req request.RegisterReq) (resp *response.RegisterResp, err error) {
	resp = new(response.RegisterResp)
	var user = &models.BlogUser{
		Email:    req.Email,
		Password: req.Password,
	}
	if resp.Token, err = Email.GetImpl(user).UserRegister(); err != nil {
		return nil, err
	}
	return
}

// LoginWithAccountHandle 邮箱登录
func (receiver *Service) LoginWithAccountHandle(c *gin.Context, req request.LoginReq) (resp *response.LoginResp, err error) {
	resp = new(response.LoginResp)
	var user = &models.BlogUser{
		Account:  req.Account,
		Password: req.Password,
	}
	if resp.Token, err = Account.GetImpl(user).UserLogin(); err != nil {
		return nil, err
	}
	return
}

// LoginWithEmailHandle 账号登录
func (receiver *Service) LoginWithEmailHandle(c *gin.Context, req request.LoginReq) (resp *response.LoginResp, err error) {
	resp = new(response.LoginResp)
	var user = &models.BlogUser{
		Email:    req.Email,
		Password: req.Password,
	}
	if resp.Token, err = Email.GetImpl(user).UserLogin(); err != nil {
		return nil, err
	}
	return
}
