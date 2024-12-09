package main

import (
	"blog/configs"
	_ "blog/init"
	"blog/internal/router"
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download
//go:generate go env -w GOOS=linux
//go:generate go env -w GOARCH=amd64
//go:generate go build -o wjy_BlogProject
func main() {
	startServer()
}

// 启动一个http服务
func startServer() {
	serverAddr := fmt.Sprintf("127.0.0.1:%s", configs.Config.System.Port)
	s := &http.Server{
		Addr:           serverAddr,
		Handler:        router.InitRouter(),
		MaxHeaderBytes: 1 << 20 * 30, // 30MB
	}
	zap.L().Info("服务启动成功，访问地址为：http://" + serverAddr)
	err := s.ListenAndServe()
	if err != nil {
		panic(err.Error())
		return
	}
}
