/*
@Time : 2024/12/9 16:22
@Author : linx
@File : auth.go
@dsc:
*/

package rear

import (
	"blog/internal/api"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
}

func NewAuthRouter() *AuthRouter {
	return &AuthRouter{}
}
func (receiver *AuthRouter) Init(group *gin.RouterGroup) {
	authRouter := group.Group("auth")
	authApi := api.NewAuthApi()
	{
		authRouter.POST("login", authApi.Login)
		authRouter.POST("register", authApi.Register)
	}
}
