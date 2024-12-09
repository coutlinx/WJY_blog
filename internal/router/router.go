/*
@Time : 2024/12/6 14:17
@Author : linx
@File : router.go
@dsc: 路由
*/

package router

import (
	"blog/configs"
	"blog/internal/middleware"
	"blog/internal/router/rear"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	if configs.Config.System.Env == "dev" {
		gin.SetMode(gin.DebugMode)
	} else if configs.Config.System.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	var RouterPrefix = "api/v1/"

	Router := gin.New()
	Router.Use(middleware.Cors()) // 放行全部跨域请求
	Router.Static("/public", "./public")
	Router.LoadHTMLGlob("templates/*")
	Router.Use(middleware.GinLogger(), gin.Recovery())

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})

		PublicGroup.GET("register", func(c *gin.Context) {
			c.HTML(http.StatusOK, "register.html", gin.H{})
		})

		PublicGroup.GET("login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", gin.H{})
		})
	}

	PrivateGroup := Router.Group(RouterPrefix)
	{
		rear.NewAuthRouter().Init(PrivateGroup)
	}
	return Router
}
