/*
@Time : 2024/12/9 16:13
@Author : linx
@File : auth.go
@dsc:
*/

// Package api 提供了与认证相关的API接口，包括用户注册和登录等功能。
// 本包通过调用内部的服务层来处理具体的业务逻辑，并返回相应的结果给客户端。
package api

import (
	"blog/internal/models/request"
	"blog/internal/models/response"
	"blog/internal/service/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthApi 是用于处理认证相关API的结构体，它包含了注册、登录等方法来对外提供服务。
type AuthApi struct {
}

// NewAuthApi 创建一个新的AuthApi实例，用于后续调用认证相关的API方法。
// 返回值：
// - *AuthApi: 返回创建好的AuthApi指针实例。
func NewAuthApi() *AuthApi {
	return &AuthApi{}
}

var service = auth.NewAuthService()

// @Summary 用户注册
// @Description 根据提供的账号或邮箱进行用户注册操作
// @Tags 认证（Auth）
// @Accept json
// @Produce json
// @Param registerReq body request.RegisterReq true "注册请求参数"
// @Success 200 {object} response.RegisterResp "注册成功返回数据"
// @Failure 400 {object}  "请求参数错误、账号或邮箱未提供等错误情况返回信息"
// @Router /api/v1/auth/register [post]

func (api *AuthApi) Register(c *gin.Context) {
	var req request.RegisterReq
	var err error
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	var resp *response.RegisterResp
	if len(req.Account) > 0 {
		if resp, err = service.RegisterWithAccountHandle(c, req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
			})
			return
		}
	} else if len(req.Email) > 0 {
		if resp, err = service.RegisterWithEmailHandle(c, req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "请输入账号或邮箱",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "注册成功",
		"data": resp,
	})
}

// @Summary 用户登录
// @Description 根据提供的账号或邮箱进行用户登录操作
// @Tags 认证（Auth）
// @Accept json
// @Produce json
// @Param loginReq body request.LoginReq true "登录请求参数"
// @Success 200 {object} response.LoginResp "登录成功返回数据"
// @Failure 400 {object}  "请求参数错误、账号或邮箱未提供等错误情况返回信息"
// @Router /api/v1/auth/login [post]

func (api *AuthApi) Login(c *gin.Context) {
	var req request.LoginReq
	var err error
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}

	var resp *response.LoginResp
	if len(req.Account) > 0 {
		if resp, err = service.LoginWithAccountHandle(c, req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
			})
			return
		}
	} else if len(req.Email) > 0 {
		if resp, err = service.LoginWithEmailHandle(c, req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "请输入账号或邮箱",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "登录成功",
		"data": resp,
	})
}
