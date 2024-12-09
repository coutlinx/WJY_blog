/*
@Time : 2024/11/13 16:55
@Author : linx
@File : global.go
@dsc: 系统的全局参数
*/

package global

import (
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	Redis *redis.Pool
)
