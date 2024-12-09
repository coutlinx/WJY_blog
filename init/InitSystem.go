/*
@Time : 2024/11/13 15:49
@Author : linx
@File : InitSystem.go
@dsc: 初始化系统信息
*/

package init

import (
	"blog/configs"
	"go.uber.org/zap"
)

func initSystem() {
	zap.S().Infof("env:[%s]\tname:%s", configs.Config.System.Env, configs.Config.System.Name)
}
