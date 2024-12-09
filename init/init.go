/*
@Time : 2024/11/13 15:10
@Author : linx
@File : init.go
@dsc: 用于执行所有的初始化操作
*/

package init

import (
	"blog/configs"
	"blog/internal/constant"
	"blog/internal/models"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func init() {
	// 打开 YAML 文件
	file, err := os.Open(constant.ConfigPath)
	if err != nil {
		panic(fmt.Sprintf("读取配置文件失败,err:%s", err.Error()))
	}

	defer func() {
		_ = file.Close()
	}()

	// 创建解析器
	decoder := yaml.NewDecoder(file)

	// 解析 YAML 数据
	err = decoder.Decode(&configs.Config)
	if err != nil {
		panic(fmt.Sprintf("解析配置文件失败,err:%s", err.Error()))
		return
	}
	initLogger() // 初始化日志
	initSystem() // 初始化系统

	initDatabase()

	err = models.AutoMigrate() // 自动迁移数据库
	if err != nil {
		panic(fmt.Sprintf("自动迁移数据库失败,err:%s", err.Error()))
	}
}
